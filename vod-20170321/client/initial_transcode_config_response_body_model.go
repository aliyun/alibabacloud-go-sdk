// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitialTranscodeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitialTranscodeConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *InitialTranscodeConfigResponseBody
	GetResult() *bool
}

type InitialTranscodeConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InitialTranscodeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitialTranscodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *InitialTranscodeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitialTranscodeConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InitialTranscodeConfigResponseBody) SetRequestId(v string) *InitialTranscodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitialTranscodeConfigResponseBody) SetResult(v bool) *InitialTranscodeConfigResponseBody {
	s.Result = &v
	return s
}

func (s *InitialTranscodeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
