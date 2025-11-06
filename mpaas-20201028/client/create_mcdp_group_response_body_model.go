// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMcdpGroupResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcdpGroupResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMcdpGroupResponseBodyResultContent) *CreateMcdpGroupResponseBody
	GetResultContent() *CreateMcdpGroupResponseBodyResultContent
	SetResultMessage(v string) *CreateMcdpGroupResponseBody
	GetResultMessage() *string
}

type CreateMcdpGroupResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMcdpGroupResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcdpGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcdpGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcdpGroupResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcdpGroupResponseBody) GetResultContent() *CreateMcdpGroupResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMcdpGroupResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcdpGroupResponseBody) SetRequestId(v string) *CreateMcdpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcdpGroupResponseBody) SetResultCode(v string) *CreateMcdpGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcdpGroupResponseBody) SetResultContent(v *CreateMcdpGroupResponseBodyResultContent) *CreateMcdpGroupResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMcdpGroupResponseBody) SetResultMessage(v string) *CreateMcdpGroupResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcdpGroupResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcdpGroupResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcdpGroupResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpGroupResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMcdpGroupResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *CreateMcdpGroupResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *CreateMcdpGroupResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *CreateMcdpGroupResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcdpGroupResponseBodyResultContent) SetCode(v string) *CreateMcdpGroupResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *CreateMcdpGroupResponseBodyResultContent) SetData(v string) *CreateMcdpGroupResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *CreateMcdpGroupResponseBodyResultContent) SetMessage(v string) *CreateMcdpGroupResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *CreateMcdpGroupResponseBodyResultContent) SetSuccess(v bool) *CreateMcdpGroupResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *CreateMcdpGroupResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
