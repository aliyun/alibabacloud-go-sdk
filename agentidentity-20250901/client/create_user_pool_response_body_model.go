// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserPoolResponseBody
	GetRequestId() *string
	SetUserPool(v *CreateUserPoolResponseBodyUserPool) *CreateUserPoolResponseBody
	GetUserPool() *CreateUserPoolResponseBodyUserPool
}

type CreateUserPoolResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserPool  *CreateUserPoolResponseBodyUserPool `json:"UserPool,omitempty" xml:"UserPool,omitempty" type:"Struct"`
}

func (s CreateUserPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserPoolResponseBody) GetUserPool() *CreateUserPoolResponseBodyUserPool {
	return s.UserPool
}

func (s *CreateUserPoolResponseBody) SetRequestId(v string) *CreateUserPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserPoolResponseBody) SetUserPool(v *CreateUserPoolResponseBodyUserPool) *CreateUserPoolResponseBody {
	s.UserPool = v
	return s
}

func (s *CreateUserPoolResponseBody) Validate() error {
	if s.UserPool != nil {
		if err := s.UserPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserPoolResponseBodyUserPool struct {
	// example:
	//
	// 2026-05-07T06:19:17Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2026-05-07T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// up_xxxxxxxxxxxxxxxxxxxx
	UserPoolId *string `json:"UserPoolId,omitempty" xml:"UserPoolId,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateUserPoolResponseBodyUserPool) String() string {
	return dara.Prettify(s)
}

func (s CreateUserPoolResponseBodyUserPool) GoString() string {
	return s.String()
}

func (s *CreateUserPoolResponseBodyUserPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateUserPoolResponseBodyUserPool) GetDescription() *string {
	return s.Description
}

func (s *CreateUserPoolResponseBodyUserPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateUserPoolResponseBodyUserPool) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *CreateUserPoolResponseBodyUserPool) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateUserPoolResponseBodyUserPool) SetCreateTime(v string) *CreateUserPoolResponseBodyUserPool {
	s.CreateTime = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetDescription(v string) *CreateUserPoolResponseBodyUserPool {
	s.Description = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetUpdateTime(v string) *CreateUserPoolResponseBodyUserPool {
	s.UpdateTime = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetUserPoolId(v string) *CreateUserPoolResponseBodyUserPool {
	s.UserPoolId = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) SetUserPoolName(v string) *CreateUserPoolResponseBodyUserPool {
	s.UserPoolName = &v
	return s
}

func (s *CreateUserPoolResponseBodyUserPool) Validate() error {
	return dara.Validate(s)
}
