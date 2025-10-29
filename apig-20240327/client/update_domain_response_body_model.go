// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDomainResponseBody
	GetCode() *string
	SetData(v *UpdateDomainResponseBodyData) *UpdateDomainResponseBody
	GetData() *UpdateDomainResponseBodyData
	SetMessage(v string) *UpdateDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDomainResponseBody
	GetRequestId() *string
}

type UpdateDomainResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response parameters.
	Data *UpdateDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. You can use this value to trace the API call.
	//
	// example:
	//
	// 4BACB05C-3FE2-588F-9148-700C5C026B74
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDomainResponseBody) GetData() *UpdateDomainResponseBodyData {
	return s.Data
}

func (s *UpdateDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDomainResponseBody) SetCode(v string) *UpdateDomainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDomainResponseBody) SetData(v *UpdateDomainResponseBodyData) *UpdateDomainResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDomainResponseBody) SetMessage(v string) *UpdateDomainResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDomainResponseBody) SetRequestId(v string) *UpdateDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDomainResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDomainResponseBodyData struct {
	// The released version ID.
	//
	// example:
	//
	// apr-xxx
	DeployRevisionId *string `json:"deployRevisionId,omitempty" xml:"deployRevisionId,omitempty"`
}

func (s UpdateDomainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponseBodyData) GetDeployRevisionId() *string {
	return s.DeployRevisionId
}

func (s *UpdateDomainResponseBodyData) SetDeployRevisionId(v string) *UpdateDomainResponseBodyData {
	s.DeployRevisionId = &v
	return s
}

func (s *UpdateDomainResponseBodyData) Validate() error {
	return dara.Validate(s)
}
