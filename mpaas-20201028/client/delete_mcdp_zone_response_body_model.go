// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcdpZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMcdpZoneResponseBody
	GetRequestId() *string
	SetResultCode(v string) *DeleteMcdpZoneResponseBody
	GetResultCode() *string
	SetResultContent(v *DeleteMcdpZoneResponseBodyResultContent) *DeleteMcdpZoneResponseBody
	GetResultContent() *DeleteMcdpZoneResponseBodyResultContent
	SetResultMessage(v string) *DeleteMcdpZoneResponseBody
	GetResultMessage() *string
}

type DeleteMcdpZoneResponseBody struct {
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *DeleteMcdpZoneResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s DeleteMcdpZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMcdpZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMcdpZoneResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *DeleteMcdpZoneResponseBody) GetResultContent() *DeleteMcdpZoneResponseBodyResultContent {
	return s.ResultContent
}

func (s *DeleteMcdpZoneResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *DeleteMcdpZoneResponseBody) SetRequestId(v string) *DeleteMcdpZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMcdpZoneResponseBody) SetResultCode(v string) *DeleteMcdpZoneResponseBody {
	s.ResultCode = &v
	return s
}

func (s *DeleteMcdpZoneResponseBody) SetResultContent(v *DeleteMcdpZoneResponseBodyResultContent) *DeleteMcdpZoneResponseBody {
	s.ResultContent = v
	return s
}

func (s *DeleteMcdpZoneResponseBody) SetResultMessage(v string) *DeleteMcdpZoneResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *DeleteMcdpZoneResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMcdpZoneResponseBodyResultContent struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMcdpZoneResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcdpZoneResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *DeleteMcdpZoneResponseBodyResultContent) GetCode() *string {
	return s.Code
}

func (s *DeleteMcdpZoneResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *DeleteMcdpZoneResponseBodyResultContent) GetMessage() *string {
	return s.Message
}

func (s *DeleteMcdpZoneResponseBodyResultContent) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMcdpZoneResponseBodyResultContent) SetCode(v string) *DeleteMcdpZoneResponseBodyResultContent {
	s.Code = &v
	return s
}

func (s *DeleteMcdpZoneResponseBodyResultContent) SetData(v string) *DeleteMcdpZoneResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *DeleteMcdpZoneResponseBodyResultContent) SetMessage(v string) *DeleteMcdpZoneResponseBodyResultContent {
	s.Message = &v
	return s
}

func (s *DeleteMcdpZoneResponseBodyResultContent) SetSuccess(v bool) *DeleteMcdpZoneResponseBodyResultContent {
	s.Success = &v
	return s
}

func (s *DeleteMcdpZoneResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
