// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateClientUserStatusResponseBody
	GetRequestId() *string
}

type UpdateClientUserStatusResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClientUserStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientUserStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClientUserStatusResponseBody) SetRequestId(v string) *UpdateClientUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClientUserStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
