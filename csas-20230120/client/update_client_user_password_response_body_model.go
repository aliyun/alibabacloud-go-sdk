// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateClientUserPasswordResponseBody
	GetRequestId() *string
}

type UpdateClientUserPasswordResponseBody struct {
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClientUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClientUserPasswordResponseBody) SetRequestId(v string) *UpdateClientUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClientUserPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
