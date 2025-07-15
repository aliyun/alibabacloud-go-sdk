// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveStreamTranscodeResponseBody
	GetRequestId() *string
}

type UpdateLiveStreamTranscodeResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveStreamTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveStreamTranscodeResponseBody) SetRequestId(v string) *UpdateLiveStreamTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveStreamTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
