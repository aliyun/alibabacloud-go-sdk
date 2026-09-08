// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBlacklistCallTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveBlacklistCallTaggingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveBlacklistCallTaggingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveBlacklistCallTaggingResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveBlacklistCallTaggingResponseBody
	GetRequestId() *string
}

type RemoveBlacklistCallTaggingResponseBody struct {
	// Response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request.
	//
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveBlacklistCallTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveBlacklistCallTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBlacklistCallTaggingResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveBlacklistCallTaggingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveBlacklistCallTaggingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveBlacklistCallTaggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveBlacklistCallTaggingResponseBody) SetCode(v string) *RemoveBlacklistCallTaggingResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveBlacklistCallTaggingResponseBody) SetHttpStatusCode(v int32) *RemoveBlacklistCallTaggingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveBlacklistCallTaggingResponseBody) SetMessage(v string) *RemoveBlacklistCallTaggingResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveBlacklistCallTaggingResponseBody) SetRequestId(v string) *RemoveBlacklistCallTaggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveBlacklistCallTaggingResponseBody) Validate() error {
	return dara.Validate(s)
}
