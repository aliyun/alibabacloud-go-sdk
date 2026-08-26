// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcAsrTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *StopRtcAsrTaskResponseBody
	GetDescription() *string
	SetRequestId(v string) *StopRtcAsrTaskResponseBody
	GetRequestId() *string
	SetRetCode(v int64) *StopRtcAsrTaskResponseBody
	GetRetCode() *int64
}

type StopRtcAsrTaskResponseBody struct {
	// The result of the call. The value is \\`success\\` if the call is successful. Otherwise, an error message is returned.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D208CC1-27C9-51E9-82B8-A6682D466421
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. A value of \\`2000\\` indicates that the call was successful. Other values indicate that the call failed.
	//
	// example:
	//
	// 2000
	RetCode *int64 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s StopRtcAsrTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRtcAsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopRtcAsrTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *StopRtcAsrTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRtcAsrTaskResponseBody) GetRetCode() *int64 {
	return s.RetCode
}

func (s *StopRtcAsrTaskResponseBody) SetDescription(v string) *StopRtcAsrTaskResponseBody {
	s.Description = &v
	return s
}

func (s *StopRtcAsrTaskResponseBody) SetRequestId(v string) *StopRtcAsrTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRtcAsrTaskResponseBody) SetRetCode(v int64) *StopRtcAsrTaskResponseBody {
	s.RetCode = &v
	return s
}

func (s *StopRtcAsrTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
