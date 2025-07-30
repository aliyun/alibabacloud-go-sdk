// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcountName(v string) *CreateAccountResponseBody
	GetAcountName() *string
	SetInstanceId(v string) *CreateAccountResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateAccountResponseBody
	GetRequestId() *string
}

type CreateAccountResponseBody struct {
	// The name of the account.
	//
	// example:
	//
	// demoaccount
	AcountName *string `json:"AcountName,omitempty" xml:"AcountName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABAF95F6-35C1-4177-AF3A-70969EBD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) GetAcountName() *string {
	return s.AcountName
}

func (s *CreateAccountResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccountResponseBody) SetAcountName(v string) *CreateAccountResponseBody {
	s.AcountName = &v
	return s
}

func (s *CreateAccountResponseBody) SetInstanceId(v string) *CreateAccountResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
