// Copyright 2022 Forest Urquhart. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE file.

// The MapBox Geostats project is licensed under the ISC License.
// See https://github.com/mapbox/mapbox-geostats for schema information.

// Package geostats provides a MapBox Geostats compatible geographical tile data statistics provider.
package geostats

type AttributeType int64

const (
	String AttributeType = iota
	Number
	Boolean
	Null
	Mixed
)

type Stats struct {
	LayerCount int64
	Layers     []Layer
}

type Layer struct {
	Layer          string
	Count          int64
	Geometry       string
	AttributeCount int64
	Attributes     []Attribute
}

type Attribute struct {
	Attribute string
	Count     int64
	Type      AttributeType
	Values    []interface{}
	Min       *int64
	Max       *int64
}
