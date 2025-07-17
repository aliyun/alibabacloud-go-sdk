// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
}

type ChangeResourceGroupResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 59F86F37-787A-52DB-9475-DB5A255517F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
