// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCpuHighAgentStreamResponseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CpuHighAgentStreamResponseResponseBody
	GetCode() *string
	SetData(v string) *CpuHighAgentStreamResponseResponseBody
	GetData() *string
	SetMessage(v string) *CpuHighAgentStreamResponseResponseBody
	GetMessage() *string
	SetRequestId(v string) *CpuHighAgentStreamResponseResponseBody
	GetRequestId() *string
}

type CpuHighAgentStreamResponseResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// <SSEResponse>
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CpuHighAgentStreamResponseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CpuHighAgentStreamResponseResponseBody) GoString() string {
	return s.String()
}

func (s *CpuHighAgentStreamResponseResponseBody) GetCode() *string {
	return s.Code
}

func (s *CpuHighAgentStreamResponseResponseBody) GetData() *string {
	return s.Data
}

func (s *CpuHighAgentStreamResponseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CpuHighAgentStreamResponseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CpuHighAgentStreamResponseResponseBody) SetCode(v string) *CpuHighAgentStreamResponseResponseBody {
	s.Code = &v
	return s
}

func (s *CpuHighAgentStreamResponseResponseBody) SetData(v string) *CpuHighAgentStreamResponseResponseBody {
	s.Data = &v
	return s
}

func (s *CpuHighAgentStreamResponseResponseBody) SetMessage(v string) *CpuHighAgentStreamResponseResponseBody {
	s.Message = &v
	return s
}

func (s *CpuHighAgentStreamResponseResponseBody) SetRequestId(v string) *CpuHighAgentStreamResponseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CpuHighAgentStreamResponseResponseBody) Validate() error {
	return dara.Validate(s)
}
