// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryLinkResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryLinkResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryLinkResponseBodyResultContent) *QueryLinkResponseBody
	GetResultContent() *QueryLinkResponseBodyResultContent
	SetResultMessage(v string) *QueryLinkResponseBody
	GetResultMessage() *string
}

type QueryLinkResponseBody struct {
	RequestId     *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryLinkResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLinkResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLinkResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryLinkResponseBody) GetResultContent() *QueryLinkResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryLinkResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryLinkResponseBody) SetRequestId(v string) *QueryLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLinkResponseBody) SetResultCode(v string) *QueryLinkResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryLinkResponseBody) SetResultContent(v *QueryLinkResponseBodyResultContent) *QueryLinkResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryLinkResponseBody) SetResultMessage(v string) *QueryLinkResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryLinkResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryLinkResponseBodyResultContent struct {
	Data    interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Target  *string     `json:"Target,omitempty" xml:"Target,omitempty"`
	Version *string     `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryLinkResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryLinkResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryLinkResponseBodyResultContent) GetData() interface{} {
	return s.Data
}

func (s *QueryLinkResponseBodyResultContent) GetTarget() *string {
	return s.Target
}

func (s *QueryLinkResponseBodyResultContent) GetVersion() *string {
	return s.Version
}

func (s *QueryLinkResponseBodyResultContent) SetData(v interface{}) *QueryLinkResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryLinkResponseBodyResultContent) SetTarget(v string) *QueryLinkResponseBodyResultContent {
	s.Target = &v
	return s
}

func (s *QueryLinkResponseBodyResultContent) SetVersion(v string) *QueryLinkResponseBodyResultContent {
	s.Version = &v
	return s
}

func (s *QueryLinkResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
