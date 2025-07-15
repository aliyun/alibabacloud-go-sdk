// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLiveStreamTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCustomLiveStreamTranscodeResponseBody
	GetRequestId() *string
}

type UpdateCustomLiveStreamTranscodeResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomLiveStreamTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLiveStreamTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomLiveStreamTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomLiveStreamTranscodeResponseBody) SetRequestId(v string) *UpdateCustomLiveStreamTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
