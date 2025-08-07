// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDBLinkResponseBody
	GetRequestId() *string
}

type CreateDBLinkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F9F1CB1A-B1D5-4EF5-A53A-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBLinkResponseBody) SetRequestId(v string) *CreateDBLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
