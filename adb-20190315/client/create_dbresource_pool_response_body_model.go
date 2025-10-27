// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResourcePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBResourcePoolResponseBody
	GetRequestId() *string
}

type CreateDBResourcePoolResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBResourcePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResourcePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBResourcePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBResourcePoolResponseBody) SetRequestId(v string) *CreateDBResourcePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBResourcePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
