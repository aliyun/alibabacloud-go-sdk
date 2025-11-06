// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpAimResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMcdpAimResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcdpAimResponseBody
	GetResultCode() *string
	SetResultContent(v *DeleteMcdpAimResponseBodyResultContent) *DeleteMcdpAimResponseBody
	GetResultContent() *DeleteMcdpAimResponseBodyResultContent
	SetResultMessage(v string) *DeleteMcdpAimResponseBody
	GetResultMessage() *string
}

type DeleteMcdpAimResponseBody struct {
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *DeleteMcdpAimResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMcdpAimResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpAimResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcdpAimResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcdpAimResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcdpAimResponseBody) GetResultContent() *DeleteMcdpAimResponseBodyResultContent {
	return s.ResultContent
}

func (s *DeleteMcdpAimResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcdpAimResponseBody) SetRequestId(v string) *DeleteMcdpAimResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcdpAimResponseBody) SetResultCode(v string) *DeleteMcdpAimResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcdpAimResponseBody) SetResultContent(v *DeleteMcdpAimResponseBodyResultContent) *DeleteMcdpAimResponseBody {
	s.ResultContent = v
	return s
}

func (s *DeleteMcdpAimResponseBody) SetResultMessage(v string) *DeleteMcdpAimResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcdpAimResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMcdpAimResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcdpAimResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpAimResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *DeleteMcdpAimResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *DeleteMcdpAimResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *DeleteMcdpAimResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcdpAimResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcdpAimResponseBodyResultContent) SetCode(v string) *DeleteMcdpAimResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *DeleteMcdpAimResponseBodyResultContent) SetData(v string) *DeleteMcdpAimResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *DeleteMcdpAimResponseBodyResultContent) SetMessage(v string) *DeleteMcdpAimResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *DeleteMcdpAimResponseBodyResultContent) SetSuccess(v bool) *DeleteMcdpAimResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *DeleteMcdpAimResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
