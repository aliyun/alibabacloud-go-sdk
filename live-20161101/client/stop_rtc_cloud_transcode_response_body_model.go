// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcCloudTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRtcCloudTranscodeResponseBody
	GetRequestId() *string
}

type StopRtcCloudTranscodeResponseBody struct {
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRtcCloudTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRtcCloudTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *StopRtcCloudTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRtcCloudTranscodeResponseBody) SetRequestId(v string) *StopRtcCloudTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRtcCloudTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
