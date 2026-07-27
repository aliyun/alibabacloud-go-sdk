// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryRowDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngineName(v string) *DsgQueryRowDetailRequest
	GetEngineName() *string
	SetInstId(v string) *DsgQueryRowDetailRequest
	GetInstId() *string
	SetPageNo(v int64) *DsgQueryRowDetailRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DsgQueryRowDetailRequest
	GetPageSize() *int64
}

type DsgQueryRowDetailRequest struct {
	// The engine type. Valid values:
	//
	// - ODPS.ODPS
	//
	// - EMR
	//
	// - HOLO.POSTGRES
	//
	// This parameter is required.
	//
	// example:
	//
	// EMR
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20260706102936ec393b1a03ae0d4atarget
	InstId *string `json:"InstId,omitempty" xml:"InstId,omitempty"`
	// The page number. Minimum value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DsgQueryRowDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryRowDetailRequest) GoString() string {
	return s.String()
}

func (s *DsgQueryRowDetailRequest) GetEngineName() *string {
	return s.EngineName
}

func (s *DsgQueryRowDetailRequest) GetInstId() *string {
	return s.InstId
}

func (s *DsgQueryRowDetailRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DsgQueryRowDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DsgQueryRowDetailRequest) SetEngineName(v string) *DsgQueryRowDetailRequest {
	s.EngineName = &v
	return s
}

func (s *DsgQueryRowDetailRequest) SetInstId(v string) *DsgQueryRowDetailRequest {
	s.InstId = &v
	return s
}

func (s *DsgQueryRowDetailRequest) SetPageNo(v int64) *DsgQueryRowDetailRequest {
	s.PageNo = &v
	return s
}

func (s *DsgQueryRowDetailRequest) SetPageSize(v int64) *DsgQueryRowDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DsgQueryRowDetailRequest) Validate() error {
	return dara.Validate(s)
}
