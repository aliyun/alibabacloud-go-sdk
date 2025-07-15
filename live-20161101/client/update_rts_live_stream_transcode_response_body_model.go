// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtsLiveStreamTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRtsLiveStreamTranscodeResponseBody
	GetRequestId() *string
}

type UpdateRtsLiveStreamTranscodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRtsLiveStreamTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtsLiveStreamTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRtsLiveStreamTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRtsLiveStreamTranscodeResponseBody) SetRequestId(v string) *UpdateRtsLiveStreamTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
