// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishObjectListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishObjectListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *PublishObjectListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishObjectListResponseBody
	GetMessage() *string
	SetPublishResult(v *PublishObjectListResponseBodyPublishResult) *PublishObjectListResponseBody
	GetPublishResult() *PublishObjectListResponseBodyPublishResult
	SetRequestId(v string) *PublishObjectListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishObjectListResponseBody
	GetSuccess() *bool
}

type PublishObjectListResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message       *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PublishResult *PublishObjectListResponseBodyPublishResult `json:"PublishResult,omitempty" xml:"PublishResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishObjectListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishObjectListResponseBody) GoString() string {
	return s.String()
}

func (s *PublishObjectListResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishObjectListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishObjectListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishObjectListResponseBody) GetPublishResult() *PublishObjectListResponseBodyPublishResult {
	return s.PublishResult
}

func (s *PublishObjectListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishObjectListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishObjectListResponseBody) SetCode(v string) *PublishObjectListResponseBody {
	s.Code = &v
	return s
}

func (s *PublishObjectListResponseBody) SetHttpStatusCode(v int32) *PublishObjectListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishObjectListResponseBody) SetMessage(v string) *PublishObjectListResponseBody {
	s.Message = &v
	return s
}

func (s *PublishObjectListResponseBody) SetPublishResult(v *PublishObjectListResponseBodyPublishResult) *PublishObjectListResponseBody {
	s.PublishResult = v
	return s
}

func (s *PublishObjectListResponseBody) SetRequestId(v string) *PublishObjectListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishObjectListResponseBody) SetSuccess(v bool) *PublishObjectListResponseBody {
	s.Success = &v
	return s
}

func (s *PublishObjectListResponseBody) Validate() error {
	return dara.Validate(s)
}

type PublishObjectListResponseBodyPublishResult struct {
	SubmitIdList []*int64 `json:"SubmitIdList,omitempty" xml:"SubmitIdList,omitempty" type:"Repeated"`
}

func (s PublishObjectListResponseBodyPublishResult) String() string {
	return dara.Prettify(s)
}

func (s PublishObjectListResponseBodyPublishResult) GoString() string {
	return s.String()
}

func (s *PublishObjectListResponseBodyPublishResult) GetSubmitIdList() []*int64 {
	return s.SubmitIdList
}

func (s *PublishObjectListResponseBodyPublishResult) SetSubmitIdList(v []*int64) *PublishObjectListResponseBodyPublishResult {
	s.SubmitIdList = v
	return s
}

func (s *PublishObjectListResponseBodyPublishResult) Validate() error {
	return dara.Validate(s)
}
