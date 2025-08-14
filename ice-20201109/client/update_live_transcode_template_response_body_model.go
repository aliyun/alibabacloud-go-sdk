// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveTranscodeTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveTranscodeTemplateResponseBody
	GetRequestId() *string
}

type UpdateLiveTranscodeTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveTranscodeTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveTranscodeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveTranscodeTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveTranscodeTemplateResponseBody) SetRequestId(v string) *UpdateLiveTranscodeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveTranscodeTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
