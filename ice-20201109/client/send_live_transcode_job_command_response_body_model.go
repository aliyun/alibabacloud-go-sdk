// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveTranscodeJobCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendLiveTranscodeJobCommandResponseBody
	GetRequestId() *string
}

type SendLiveTranscodeJobCommandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendLiveTranscodeJobCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendLiveTranscodeJobCommandResponseBody) GoString() string {
	return s.String()
}

func (s *SendLiveTranscodeJobCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendLiveTranscodeJobCommandResponseBody) SetRequestId(v string) *SendLiveTranscodeJobCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendLiveTranscodeJobCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
