// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindDBResourceGroupWithUserResponseBody
	GetRequestId() *string
}

type BindDBResourceGroupWithUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDBResourceGroupWithUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindDBResourceGroupWithUserResponseBody) SetRequestId(v string) *BindDBResourceGroupWithUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindDBResourceGroupWithUserResponseBody) Validate() error {
	return dara.Validate(s)
}
