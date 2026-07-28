// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResourceControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindResourceControlResponseBody
	GetRequestId() *string
}

type BindResourceControlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindResourceControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindResourceControlResponseBody) GoString() string {
	return s.String()
}

func (s *BindResourceControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindResourceControlResponseBody) SetRequestId(v string) *BindResourceControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindResourceControlResponseBody) Validate() error {
	return dara.Validate(s)
}
