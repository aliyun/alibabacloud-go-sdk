// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateMcubeNebulaResourceReslult(v *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) *CreateMcubeNebulaResourceResponseBody
	GetCreateMcubeNebulaResourceReslult() *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult
	SetRequestId(v string) *CreateMcubeNebulaResourceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcubeNebulaResourceResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *CreateMcubeNebulaResourceResponseBody
	GetResultMessage() *string
}

type CreateMcubeNebulaResourceResponseBody struct {
	CreateMcubeNebulaResourceReslult *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult `json:"CreateMcubeNebulaResourceReslult,omitempty" xml:"CreateMcubeNebulaResourceReslult,omitempty" type:"Struct"`
	RequestId                        *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                       *string                                                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage                    *string                                                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcubeNebulaResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaResourceResponseBody) GetCreateMcubeNebulaResourceReslult() *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult {
	return s.CreateMcubeNebulaResourceReslult
}

func (s *CreateMcubeNebulaResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeNebulaResourceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcubeNebulaResourceResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcubeNebulaResourceResponseBody) SetCreateMcubeNebulaResourceReslult(v *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) *CreateMcubeNebulaResourceResponseBody {
	s.CreateMcubeNebulaResourceReslult = v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBody) SetRequestId(v string) *CreateMcubeNebulaResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBody) SetResultCode(v string) *CreateMcubeNebulaResourceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBody) SetResultMessage(v string) *CreateMcubeNebulaResourceResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBody) Validate() error {
	if s.CreateMcubeNebulaResourceReslult != nil {
		if err := s.CreateMcubeNebulaResourceReslult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult struct {
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	NebulaResourceId *string `json:"NebulaResourceId,omitempty" xml:"NebulaResourceId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg        *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) GetNebulaResourceId() *string {
	return s.NebulaResourceId
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) SetErrorCode(v string) *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult {
	s.ErrorCode = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) SetNebulaResourceId(v string) *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult {
	s.NebulaResourceId = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) SetRequestId(v string) *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult {
	s.RequestId = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) SetResultMsg(v string) *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult {
	s.ResultMsg = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) SetSuccess(v bool) *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult {
	s.Success = &v
	return s
}

func (s *CreateMcubeNebulaResourceResponseBodyCreateMcubeNebulaResourceReslult) Validate() error {
	return dara.Validate(s)
}
