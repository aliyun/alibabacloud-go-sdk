// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveTranscodeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveTranscodeJobResponseBody
	GetRequestId() *string
}

type DeleteLiveTranscodeJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveTranscodeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveTranscodeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveTranscodeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveTranscodeJobResponseBody) SetRequestId(v string) *DeleteLiveTranscodeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveTranscodeJobResponseBody) Validate() error {
	return dara.Validate(s)
}
