// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *VoiceModerationCancelResponseBody
	GetCode() *int32
	SetMessage(v string) *VoiceModerationCancelResponseBody
	GetMessage() *string
	SetRequestId(v string) *VoiceModerationCancelResponseBody
	GetRequestId() *string
}

type VoiceModerationCancelResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A926AE2-4C96-573F-824F-0532960799F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationCancelResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *VoiceModerationCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VoiceModerationCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VoiceModerationCancelResponseBody) SetCode(v int32) *VoiceModerationCancelResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationCancelResponseBody) SetMessage(v string) *VoiceModerationCancelResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationCancelResponseBody) SetRequestId(v string) *VoiceModerationCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *VoiceModerationCancelResponseBody) Validate() error {
	return dara.Validate(s)
}
