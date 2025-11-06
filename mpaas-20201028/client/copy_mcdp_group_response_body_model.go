// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyMcdpGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyMcdpGroupResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CopyMcdpGroupResponseBody
	GetResultCode() *string
	SetResultContent(v *CopyMcdpGroupResponseBodyResultContent) *CopyMcdpGroupResponseBody
	GetResultContent() *CopyMcdpGroupResponseBodyResultContent
	SetResultMessage(v string) *CopyMcdpGroupResponseBody
	GetResultMessage() *string
}

type CopyMcdpGroupResponseBody struct {
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CopyMcdpGroupResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CopyMcdpGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyMcdpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CopyMcdpGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyMcdpGroupResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CopyMcdpGroupResponseBody) GetResultContent() *CopyMcdpGroupResponseBodyResultContent {
	return s.ResultContent
}

func (s *CopyMcdpGroupResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CopyMcdpGroupResponseBody) SetRequestId(v string) *CopyMcdpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyMcdpGroupResponseBody) SetResultCode(v string) *CopyMcdpGroupResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CopyMcdpGroupResponseBody) SetResultContent(v *CopyMcdpGroupResponseBodyResultContent) *CopyMcdpGroupResponseBody {
	s.ResultContent = v
	return s
}

func (s *CopyMcdpGroupResponseBody) SetResultMessage(v string) *CopyMcdpGroupResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CopyMcdpGroupResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyMcdpGroupResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CopyMcdpGroupResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CopyMcdpGroupResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CopyMcdpGroupResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *CopyMcdpGroupResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *CopyMcdpGroupResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *CopyMcdpGroupResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *CopyMcdpGroupResponseBodyResultContent) SetCode(v string) *CopyMcdpGroupResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *CopyMcdpGroupResponseBodyResultContent) SetData(v string) *CopyMcdpGroupResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *CopyMcdpGroupResponseBodyResultContent) SetMessage(v string) *CopyMcdpGroupResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *CopyMcdpGroupResponseBodyResultContent) SetSuccess(v bool) *CopyMcdpGroupResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *CopyMcdpGroupResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
