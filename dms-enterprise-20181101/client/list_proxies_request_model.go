// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTid(v int64) *ListProxiesRequest
	GetTid() *int64
}

type ListProxiesRequest struct {
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListProxiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProxiesRequest) GoString() string {
	return s.String()
}

func (s *ListProxiesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListProxiesRequest) SetTid(v int64) *ListProxiesRequest {
	s.Tid = &v
	return s
}

func (s *ListProxiesRequest) Validate() error {
	return dara.Validate(s)
}
