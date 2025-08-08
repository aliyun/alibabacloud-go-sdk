// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLinkResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreateLinkResponseBody
	GetResultCode() *string
	SetResultContent(v *CreateLinkResponseBodyResultContent) *CreateLinkResponseBody
	GetResultContent() *CreateLinkResponseBodyResultContent
	SetResultMessage(v string) *CreateLinkResponseBody
	GetResultMessage() *string
}

type CreateLinkResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0CC8A9CB-9BA3-13FD-A404-6E2E7461881A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ResultCode    *string                              `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *CreateLinkResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s CreateLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLinkResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreateLinkResponseBody) GetResultContent() *CreateLinkResponseBodyResultContent {
	return s.ResultContent
}

func (s *CreateLinkResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *CreateLinkResponseBody) SetRequestId(v string) *CreateLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLinkResponseBody) SetResultCode(v string) *CreateLinkResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreateLinkResponseBody) SetResultContent(v *CreateLinkResponseBodyResultContent) *CreateLinkResponseBody {
	s.ResultContent = v
	return s
}

func (s *CreateLinkResponseBody) SetResultMessage(v string) *CreateLinkResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *CreateLinkResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateLinkResponseBodyResultContent struct {
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
	// 0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateLinkResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s CreateLinkResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *CreateLinkResponseBodyResultContent) GetData() *string {
	return s.Data
}

func (s *CreateLinkResponseBodyResultContent) GetTarget() *string {
	return s.Target
}

func (s *CreateLinkResponseBodyResultContent) GetVersion() *string {
	return s.Version
}

func (s *CreateLinkResponseBodyResultContent) SetData(v string) *CreateLinkResponseBodyResultContent {
	s.Data = &v
	return s
}

func (s *CreateLinkResponseBodyResultContent) SetTarget(v string) *CreateLinkResponseBodyResultContent {
	s.Target = &v
	return s
}

func (s *CreateLinkResponseBodyResultContent) SetVersion(v string) *CreateLinkResponseBodyResultContent {
	s.Version = &v
	return s
}

func (s *CreateLinkResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}
