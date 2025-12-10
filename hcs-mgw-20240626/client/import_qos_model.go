// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportQos interface {
	dara.Model
	String() string
	GoString() string
	SetMaxBandWidth(v int64) *ImportQos
	GetMaxBandWidth() *int64
	SetMaxImportTaskQps(v int64) *ImportQos
	GetMaxImportTaskQps() *int64
}

type ImportQos struct {
	// example:
	//
	// 1073741824
	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" xml:"MaxBandWidth,omitempty"`
	// example:
	//
	// 1000
	MaxImportTaskQps *int64 `json:"MaxImportTaskQps,omitempty" xml:"MaxImportTaskQps,omitempty"`
}

func (s ImportQos) String() string {
	return dara.Prettify(s)
}

func (s ImportQos) GoString() string {
	return s.String()
}

func (s *ImportQos) GetMaxBandWidth() *int64 {
	return s.MaxBandWidth
}

func (s *ImportQos) GetMaxImportTaskQps() *int64 {
	return s.MaxImportTaskQps
}

func (s *ImportQos) SetMaxBandWidth(v int64) *ImportQos {
	s.MaxBandWidth = &v
	return s
}

func (s *ImportQos) SetMaxImportTaskQps(v int64) *ImportQos {
	s.MaxImportTaskQps = &v
	return s
}

func (s *ImportQos) Validate() error {
	return dara.Validate(s)
}
