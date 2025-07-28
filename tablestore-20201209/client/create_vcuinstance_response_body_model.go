// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVCUInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVCUInstanceResponseBody
	GetCode() *string
	SetInstanceName(v string) *CreateVCUInstanceResponseBody
	GetInstanceName() *string
	SetMessage(v string) *CreateVCUInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVCUInstanceResponseBody
	GetRequestId() *string
}

type CreateVCUInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// request id
	//
	// example:
	//
	// 39871ED2-62C0-578F-A32E-B88072D5582F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVCUInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVCUInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVCUInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVCUInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateVCUInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVCUInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVCUInstanceResponseBody) SetCode(v string) *CreateVCUInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVCUInstanceResponseBody) SetInstanceName(v string) *CreateVCUInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *CreateVCUInstanceResponseBody) SetMessage(v string) *CreateVCUInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVCUInstanceResponseBody) SetRequestId(v string) *CreateVCUInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVCUInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
