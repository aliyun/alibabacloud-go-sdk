// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *VideoModerationCancelResponseBody
	GetCode() *int32
	SetMessage(v string) *VideoModerationCancelResponseBody
	GetMessage() *string
	SetRequestId(v string) *VideoModerationCancelResponseBody
	GetRequestId() *string
}

type VideoModerationCancelResponseBody struct {
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
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6CF2815C-****-****-B52E-FF6E2****492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoModerationCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationCancelResponseBody) GoString() string {
	return s.String()
}

func (s *VideoModerationCancelResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *VideoModerationCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VideoModerationCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VideoModerationCancelResponseBody) SetCode(v int32) *VideoModerationCancelResponseBody {
	s.Code = &v
	return s
}

func (s *VideoModerationCancelResponseBody) SetMessage(v string) *VideoModerationCancelResponseBody {
	s.Message = &v
	return s
}

func (s *VideoModerationCancelResponseBody) SetRequestId(v string) *VideoModerationCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *VideoModerationCancelResponseBody) Validate() error {
	return dara.Validate(s)
}
