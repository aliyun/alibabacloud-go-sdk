// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlacklistCallTaggingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddBlacklistCallTaggingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddBlacklistCallTaggingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddBlacklistCallTaggingResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddBlacklistCallTaggingResponseBody
	GetRequestId() *string
}

type AddBlacklistCallTaggingResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9FBA26B0-462B-4D77-B78F-AF35560DBC71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBlacklistCallTaggingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBlacklistCallTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *AddBlacklistCallTaggingResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddBlacklistCallTaggingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddBlacklistCallTaggingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddBlacklistCallTaggingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBlacklistCallTaggingResponseBody) SetCode(v string) *AddBlacklistCallTaggingResponseBody {
	s.Code = &v
	return s
}

func (s *AddBlacklistCallTaggingResponseBody) SetHttpStatusCode(v int32) *AddBlacklistCallTaggingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddBlacklistCallTaggingResponseBody) SetMessage(v string) *AddBlacklistCallTaggingResponseBody {
	s.Message = &v
	return s
}

func (s *AddBlacklistCallTaggingResponseBody) SetRequestId(v string) *AddBlacklistCallTaggingResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBlacklistCallTaggingResponseBody) Validate() error {
	return dara.Validate(s)
}
