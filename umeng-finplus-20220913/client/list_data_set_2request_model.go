// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSet2Request interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v int64) *ListDataSet2Request
	GetClientId() *int64
	SetPageNo(v int32) *ListDataSet2Request
	GetPageNo() *int32
	SetPageSize(v int32) *ListDataSet2Request
	GetPageSize() *int32
}

type ListDataSet2Request struct {
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	PageNo   *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDataSet2Request) String() string {
	return dara.Prettify(s)
}

func (s ListDataSet2Request) GoString() string {
	return s.String()
}

func (s *ListDataSet2Request) GetClientId() *int64 {
	return s.ClientId
}

func (s *ListDataSet2Request) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListDataSet2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSet2Request) SetClientId(v int64) *ListDataSet2Request {
	s.ClientId = &v
	return s
}

func (s *ListDataSet2Request) SetPageNo(v int32) *ListDataSet2Request {
	s.PageNo = &v
	return s
}

func (s *ListDataSet2Request) SetPageSize(v int32) *ListDataSet2Request {
	s.PageSize = &v
	return s
}

func (s *ListDataSet2Request) Validate() error {
	return dara.Validate(s)
}
