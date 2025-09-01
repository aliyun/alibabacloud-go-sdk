// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *GetListRecordRequest
	GetCurrent() *int64
	SetPageSize(v int64) *GetListRecordRequest
	GetPageSize() *int64
	SetRegion(v string) *GetListRecordRequest
	GetRegion() *string
}

type GetListRecordRequest struct {
	// example:
	//
	// 5
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 10
	PageSize *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s GetListRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListRecordRequest) GoString() string {
	return s.String()
}

func (s *GetListRecordRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *GetListRecordRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetListRecordRequest) GetRegion() *string {
	return s.Region
}

func (s *GetListRecordRequest) SetCurrent(v int64) *GetListRecordRequest {
	s.Current = &v
	return s
}

func (s *GetListRecordRequest) SetPageSize(v int64) *GetListRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetListRecordRequest) SetRegion(v string) *GetListRecordRequest {
	s.Region = &v
	return s
}

func (s *GetListRecordRequest) Validate() error {
	return dara.Validate(s)
}
