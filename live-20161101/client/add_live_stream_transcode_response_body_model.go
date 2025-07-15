// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveStreamTranscodeResponseBody
	GetRequestId() *string
}

type AddLiveStreamTranscodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveStreamTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveStreamTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveStreamTranscodeResponseBody) SetRequestId(v string) *AddLiveStreamTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveStreamTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
