// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLinkResponseBody
	GetRequestId() *string
	SetResultCode(v string) *UpdateLinkResponseBody
	GetResultCode() *string
	SetResultContent(v *UpdateLinkResponseBodyResultContent) *UpdateLinkResponseBody
	GetResultContent() *UpdateLinkResponseBodyResultContent
	SetResultMessage(v string) *UpdateLinkResponseBody
	GetResultMessage() *string
}

type UpdateLinkResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ResultCode    *string                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *UpdateLinkResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s UpdateLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLinkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLinkResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *UpdateLinkResponseBody) GetResultContent() *UpdateLinkResponseBodyResultContent {
	return s.ResultContent
}

func (s *UpdateLinkResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *UpdateLinkResponseBody) SetRequestId(v string) *UpdateLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLinkResponseBody) SetResultCode(v string) *UpdateLinkResponseBody {
	s.ResultCode = &v
	return s
}

func (s *UpdateLinkResponseBody) SetResultContent(v *UpdateLinkResponseBodyResultContent) *UpdateLinkResponseBody {
	s.ResultContent = v
	return s
}

func (s *UpdateLinkResponseBody) SetResultMessage(v string) *UpdateLinkResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *UpdateLinkResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLinkResponseBodyResultContent struct {
	// example:
	//
	// https://xxx/xxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// https://xxx/xxx/xxx
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateLinkResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateLinkResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *UpdateLinkResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *UpdateLinkResponseBodyResultContent) GetTarget() *string {
	return s.Target
}

func (s *UpdateLinkResponseBodyResultContent) GetVersion() *string {
	return s.Version
}

func (s *UpdateLinkResponseBodyResultContent) SetData(v string) *UpdateLinkResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *UpdateLinkResponseBodyResultContent) SetTarget(v string) *UpdateLinkResponseBodyResultContent {
	s.Target = &v
	return s
}

func (s *UpdateLinkResponseBodyResultContent) SetVersion(v string) *UpdateLinkResponseBodyResultContent {
	s.Version = &v
	return s
}

func (s *UpdateLinkResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
