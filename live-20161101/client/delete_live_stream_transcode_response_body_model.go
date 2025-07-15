// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveStreamTranscodeResponseBody
	GetRequestId() *string
}

type DeleteLiveStreamTranscodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveStreamTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamTranscodeResponseBody) SetRequestId(v string) *DeleteLiveStreamTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
