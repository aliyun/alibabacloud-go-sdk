// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLiveStreamTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCustomLiveStreamTranscodeResponseBody
	GetRequestId() *string
}

type AddCustomLiveStreamTranscodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCustomLiveStreamTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLiveStreamTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomLiveStreamTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomLiveStreamTranscodeResponseBody) SetRequestId(v string) *AddCustomLiveStreamTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
