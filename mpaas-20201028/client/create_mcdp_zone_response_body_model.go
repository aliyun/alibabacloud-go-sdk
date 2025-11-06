// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcdpZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMcdpZoneResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateMcdpZoneResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateMcdpZoneResponseBodyResultContent) *CreateMcdpZoneResponseBody
	GetResultContent() *CreateMcdpZoneResponseBodyResultContent
	SetResultMessage(v string) *CreateMcdpZoneResponseBody
	GetResultMessage() *string
}

type CreateMcdpZoneResponseBody struct {
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateMcdpZoneResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateMcdpZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpZoneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcdpZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcdpZoneResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateMcdpZoneResponseBody) GetResultContent() *CreateMcdpZoneResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateMcdpZoneResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateMcdpZoneResponseBody) SetRequestId(v string) *CreateMcdpZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcdpZoneResponseBody) SetResultCode(v string) *CreateMcdpZoneResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateMcdpZoneResponseBody) SetResultContent(v *CreateMcdpZoneResponseBodyResultContent) *CreateMcdpZoneResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateMcdpZoneResponseBody) SetResultMessage(v string) *CreateMcdpZoneResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateMcdpZoneResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcdpZoneResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMcdpZoneResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateMcdpZoneResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateMcdpZoneResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *CreateMcdpZoneResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *CreateMcdpZoneResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *CreateMcdpZoneResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMcdpZoneResponseBodyResultContent) SetCode(v string) *CreateMcdpZoneResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *CreateMcdpZoneResponseBodyResultContent) SetData(v string) *CreateMcdpZoneResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *CreateMcdpZoneResponseBodyResultContent) SetMessage(v string) *CreateMcdpZoneResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *CreateMcdpZoneResponseBodyResultContent) SetSuccess(v bool) *CreateMcdpZoneResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *CreateMcdpZoneResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
