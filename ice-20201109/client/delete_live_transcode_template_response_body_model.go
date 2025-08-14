// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveTranscodeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveTranscodeTemplateResponseBody
	GetRequestId() *string
}

type DeleteLiveTranscodeTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveTranscodeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveTranscodeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveTranscodeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveTranscodeTemplateResponseBody) SetRequestId(v string) *DeleteLiveTranscodeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveTranscodeTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
