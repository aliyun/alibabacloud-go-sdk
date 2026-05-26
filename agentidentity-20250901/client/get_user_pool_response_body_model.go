// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserPoolResponseBody
	GetRequestId() *string
	SetUserPool(v *GetUserPoolResponseBodyUserPool) *GetUserPoolResponseBody
	GetUserPool() *GetUserPoolResponseBodyUserPool
}

type GetUserPoolResponseBody struct {
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserPool  *GetUserPoolResponseBodyUserPool `json:"UserPool,omitempty" xml:"UserPool,omitempty" type:"Struct"`
}

func (s GetUserPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserPoolResponseBody) GetUserPool() *GetUserPoolResponseBodyUserPool {
	return s.UserPool
}

func (s *GetUserPoolResponseBody) SetRequestId(v string) *GetUserPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserPoolResponseBody) SetUserPool(v *GetUserPoolResponseBodyUserPool) *GetUserPoolResponseBody {
	s.UserPool = v
	return s
}

func (s *GetUserPoolResponseBody) Validate() error {
	if s.UserPool != nil {
		if err := s.UserPool.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserPoolResponseBodyUserPool struct {
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

func (s GetUserPoolResponseBodyUserPool) String() string {
	return dara.Prettify(s)
}

func (s GetUserPoolResponseBodyUserPool) GoString() string {
	return s.String()
}

func (s *GetUserPoolResponseBodyUserPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserPoolResponseBodyUserPool) GetDescription() *string {
	return s.Description
}

func (s *GetUserPoolResponseBodyUserPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserPoolResponseBodyUserPool) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *GetUserPoolResponseBodyUserPool) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetUserPoolResponseBodyUserPool) SetCreateTime(v string) *GetUserPoolResponseBodyUserPool {
	s.CreateTime = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetDescription(v string) *GetUserPoolResponseBodyUserPool {
	s.Description = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetUpdateTime(v string) *GetUserPoolResponseBodyUserPool {
	s.UpdateTime = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetUserPoolId(v string) *GetUserPoolResponseBodyUserPool {
	s.UserPoolId = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) SetUserPoolName(v string) *GetUserPoolResponseBodyUserPool {
	s.UserPoolName = &v
	return s
}

func (s *GetUserPoolResponseBodyUserPool) Validate() error {
	return dara.Validate(s)
}
