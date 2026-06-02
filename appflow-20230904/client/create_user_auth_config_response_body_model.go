// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserAuthConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateUserAuthConfigResponseBody
	GetData() *string
	SetRequestId(v string) *CreateUserAuthConfigResponseBody
	GetRequestId() *string
}

type CreateUserAuthConfigResponseBody struct {
	// example:
	//
	// uac-a2253c40486c40c1b910
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 249048A1-7FF7-5D2E-A322-695420112094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserAuthConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserAuthConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserAuthConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateUserAuthConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserAuthConfigResponseBody) SetData(v string) *CreateUserAuthConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateUserAuthConfigResponseBody) SetRequestId(v string) *CreateUserAuthConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserAuthConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
