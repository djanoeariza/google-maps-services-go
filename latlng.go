// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maps

import (
	"fmt"
	"math"
)

// LatLng represents a location on the Earth.
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (l *LatLng) String() string {
	return fmt.Sprintf("%g,%g", l.Lat, l.Lng)
}

// AlmostEqual returns whether this LatLng is almost equal (below epsilon) to
// the other LatLng.
func (l *LatLng) AlmostEqual(other *LatLng, epsilon float64) bool {
	return math.Abs(l.Lat-other.Lat) < epsilon && math.Abs(l.Lng-other.Lng) < epsilon
}

// LatLngBounds represents a bounded square area on the Earth.
type LatLngBounds struct {
	NorthEast LatLng `json:"northeast"`
	SouthWest LatLng `json:"southwest"`
}

func (b *LatLngBounds) String() string {
	return fmt.Sprintf("%s|%s", b.SouthWest.String(), b.NorthEast.String())
}
