// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPagedInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v int64) *GetPagedInstanceRequest
	GetOrderId() *int64
	SetTid(v int64) *GetPagedInstanceRequest
	GetTid() *int64
}

type GetPagedInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 868*****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetPagedInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPagedInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetPagedInstanceRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetPagedInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetPagedInstanceRequest) SetOrderId(v int64) *GetPagedInstanceRequest {
	s.OrderId = &v
	return s
}

func (s *GetPagedInstanceRequest) SetTid(v int64) *GetPagedInstanceRequest {
	s.Tid = &v
	return s
}

func (s *GetPagedInstanceRequest) Validate() error {
	return dara.Validate(s)
}
