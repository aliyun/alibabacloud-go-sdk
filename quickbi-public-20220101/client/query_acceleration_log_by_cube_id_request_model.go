// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccelerationLogByCubeIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCubeId(v string) *QueryAccelerationLogByCubeIdRequest
	GetCubeId() *string
	SetEndDate(v string) *QueryAccelerationLogByCubeIdRequest
	GetEndDate() *string
	SetPageNo(v int32) *QueryAccelerationLogByCubeIdRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryAccelerationLogByCubeIdRequest
	GetPageSize() *int32
	SetStartDate(v string) *QueryAccelerationLogByCubeIdRequest
	GetStartDate() *string
}

type QueryAccelerationLogByCubeIdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7c7223ae-****-3c744528014b
	CubeId *string `json:"CubeId,omitempty" xml:"CubeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-05-15 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-04-15 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryAccelerationLogByCubeIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccelerationLogByCubeIdRequest) GoString() string {
	return s.String()
}

func (s *QueryAccelerationLogByCubeIdRequest) GetCubeId() *string {
	return s.CubeId
}

func (s *QueryAccelerationLogByCubeIdRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryAccelerationLogByCubeIdRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryAccelerationLogByCubeIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAccelerationLogByCubeIdRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryAccelerationLogByCubeIdRequest) SetCubeId(v string) *QueryAccelerationLogByCubeIdRequest {
	s.CubeId = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdRequest) SetEndDate(v string) *QueryAccelerationLogByCubeIdRequest {
	s.EndDate = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdRequest) SetPageNo(v int32) *QueryAccelerationLogByCubeIdRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdRequest) SetPageSize(v int32) *QueryAccelerationLogByCubeIdRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdRequest) SetStartDate(v string) *QueryAccelerationLogByCubeIdRequest {
	s.StartDate = &v
	return s
}

func (s *QueryAccelerationLogByCubeIdRequest) Validate() error {
	return dara.Validate(s)
}
