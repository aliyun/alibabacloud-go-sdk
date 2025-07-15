// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRtsLiveStreamTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddRtsLiveStreamTranscodeResponseBody
	GetRequestId() *string
}

type AddRtsLiveStreamTranscodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddRtsLiveStreamTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRtsLiveStreamTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *AddRtsLiveStreamTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRtsLiveStreamTranscodeResponseBody) SetRequestId(v string) *AddRtsLiveStreamTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRtsLiveStreamTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
