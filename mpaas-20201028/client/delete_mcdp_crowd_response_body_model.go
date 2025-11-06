// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpCrowdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMcdpCrowdResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcdpCrowdResponseBody
	GetResultCode() *string
	SetResultContent(v *DeleteMcdpCrowdResponseBodyResultContent) *DeleteMcdpCrowdResponseBody
	GetResultContent() *DeleteMcdpCrowdResponseBodyResultContent
	SetResultMessage(v string) *DeleteMcdpCrowdResponseBody
	GetResultMessage() *string
}

type DeleteMcdpCrowdResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                   `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *DeleteMcdpCrowdResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                   `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMcdpCrowdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcdpCrowdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcdpCrowdResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcdpCrowdResponseBody) GetResultContent() *DeleteMcdpCrowdResponseBodyResultContent {
	return s.ResultContent
}

func (s *DeleteMcdpCrowdResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcdpCrowdResponseBody) SetRequestId(v string) *DeleteMcdpCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcdpCrowdResponseBody) SetResultCode(v string) *DeleteMcdpCrowdResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcdpCrowdResponseBody) SetResultContent(v *DeleteMcdpCrowdResponseBodyResultContent) *DeleteMcdpCrowdResponseBody {
	s.ResultContent = v
	return s
}

func (s *DeleteMcdpCrowdResponseBody) SetResultMessage(v string) *DeleteMcdpCrowdResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcdpCrowdResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMcdpCrowdResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcdpCrowdResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpCrowdResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) SetCode(v string) *DeleteMcdpCrowdResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) SetData(v string) *DeleteMcdpCrowdResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) SetMessage(v string) *DeleteMcdpCrowdResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) SetSuccess(v bool) *DeleteMcdpCrowdResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *DeleteMcdpCrowdResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
