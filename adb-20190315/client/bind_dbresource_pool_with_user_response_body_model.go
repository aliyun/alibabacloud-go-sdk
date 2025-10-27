// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourcePoolWithUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindDBResourcePoolWithUserResponseBody
	GetRequestId() *string
}

type BindDBResourcePoolWithUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDBResourcePoolWithUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourcePoolWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindDBResourcePoolWithUserResponseBody) SetRequestId(v string) *BindDBResourcePoolWithUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindDBResourcePoolWithUserResponseBody) Validate() error {
	return dara.Validate(s)
}
