// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceResponseBody
	GetCode() *string
	SetInstanceId(v string) *CreateInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// InstanceAlreadyExistsErrorProblem
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// featureStore-cn-7mz2xfu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
