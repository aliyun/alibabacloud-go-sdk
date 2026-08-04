// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayModeControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PlayModeControlResponseBody
	GetCode() *int32
	SetMessage(v string) *PlayModeControlResponseBody
	GetMessage() *string
	SetRequestId(v string) *PlayModeControlResponseBody
	GetRequestId() *string
	SetResult(v *PlayModeControlResponseBodyResult) *PlayModeControlResponseBody
	GetResult() *PlayModeControlResponseBodyResult
	SetSuccess(v string) *PlayModeControlResponseBody
	GetSuccess() *string
}

type PlayModeControlResponseBody struct {
	// Return code of the invocation
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Additional information, typically used to briefly describe a failed invocation to help the caller troubleshoot the issue.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10002398812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Actual return result of the service
	Result *PlayModeControlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the invocation succeeded. true indicates success, and false indicates failure. When the value is false, check the Message field for details.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PlayModeControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlResponseBody) GoString() string {
	return s.String()
}

func (s *PlayModeControlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PlayModeControlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PlayModeControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PlayModeControlResponseBody) GetResult() *PlayModeControlResponseBodyResult {
	return s.Result
}

func (s *PlayModeControlResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PlayModeControlResponseBody) SetCode(v int32) *PlayModeControlResponseBody {
	s.Code = &v
	return s
}

func (s *PlayModeControlResponseBody) SetMessage(v string) *PlayModeControlResponseBody {
	s.Message = &v
	return s
}

func (s *PlayModeControlResponseBody) SetRequestId(v string) *PlayModeControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *PlayModeControlResponseBody) SetResult(v *PlayModeControlResponseBodyResult) *PlayModeControlResponseBody {
	s.Result = v
	return s
}

func (s *PlayModeControlResponseBody) SetSuccess(v string) *PlayModeControlResponseBody {
	s.Success = &v
	return s
}

func (s *PlayModeControlResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PlayModeControlResponseBodyResult struct {
	// Playback mode
	//
	// List loop: Repeat; Shuffle: Shuffle; Single-track loop: RepeatOne; NAT mode: Normal;
	//
	// example:
	//
	// Normal
	OpenPlayMode *string `json:"OpenPlayMode,omitempty" xml:"OpenPlayMode,omitempty"`
}

func (s PlayModeControlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PlayModeControlResponseBodyResult) GetOpenPlayMode() *string {
	return s.OpenPlayMode
}

func (s *PlayModeControlResponseBodyResult) SetOpenPlayMode(v string) *PlayModeControlResponseBodyResult {
	s.OpenPlayMode = &v
	return s
}

func (s *PlayModeControlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
