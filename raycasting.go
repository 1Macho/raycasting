package raycasting

import (
  "github.com/1Macho/geometry"
  "math"
)

func (this Scene) Raycast (ray geometry.Ray) []geometry.Point {
  var results []geometry.Point
  for _, target := range this.Targets {
    thisResult := geometry.InterceptRaySegment(ray, target)
    if (thisResult.X != math.Inf(1) && thisResult.Y != math.Inf(1)) {
      results = append(results, thisResult)
    }
  }
  return results
}

func (this Scene) ClosestRaycast (ray geometry.Ray) (bool, geometry.Point) {
  raycastHits := this.Raycast(ray)
  closestPoint := geometry.Point{math.Inf(1),math.Inf(1)}
  closestPointDistance := math.Inf(1)
  gotHit := false
  if (len(raycastHits) > 0) {
    closestPoint = raycastHits[0]
    closestPointDistance = ray.Line.Origin.Distance(closestPoint)
    gotHit = true
  }
  for _, hit := range raycastHits {
    distanceToHit := ray.Line.Origin.Distance(hit)
    if (distanceToHit < closestPointDistance) {
      closestPoint = hit
      closestPointDistance = distanceToHit
    }
  }
  return gotHit, closestPoint
}
