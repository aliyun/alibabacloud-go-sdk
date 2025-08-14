// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMediaLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopMediaLiveChannelResponseBody
	GetRequestId() *string
}

type StopMediaLiveChannelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMediaLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopMediaLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *StopMediaLiveChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopMediaLiveChannelResponseBody) SetRequestId(v string) *StopMediaLiveChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMediaLiveChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
