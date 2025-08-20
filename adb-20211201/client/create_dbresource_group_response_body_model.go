// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBResourceGroupResponseBody
	GetRequestId() *string
}

type CreateDBResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A94B6C02-7BD4-5D67-9776-3AC8317E8DD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBResourceGroupResponseBody) SetRequestId(v string) *CreateDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
