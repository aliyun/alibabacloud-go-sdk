// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateClientUserResponseBody
	GetRequestId() *string
}

type UpdateClientUserResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClientUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClientUserResponseBody) SetRequestId(v string) *UpdateClientUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClientUserResponseBody) Validate() error {
	return dara.Validate(s)
}
