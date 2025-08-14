// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveTranscodeJobResponseBody
	GetRequestId() *string
}

type UpdateLiveTranscodeJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveTranscodeJobResponseBody) SetRequestId(v string) *UpdateLiveTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
