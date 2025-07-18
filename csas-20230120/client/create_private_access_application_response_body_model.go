// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreatePrivateAccessApplicationResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *CreatePrivateAccessApplicationResponseBody
	GetRequestId() *string
}

type CreatePrivateAccessApplicationResponseBody struct {
	// The ID of the office application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrivateAccessApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePrivateAccessApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrivateAccessApplicationResponseBody) SetApplicationId(v string) *CreatePrivateAccessApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreatePrivateAccessApplicationResponseBody) SetRequestId(v string) *CreatePrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivateAccessApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
