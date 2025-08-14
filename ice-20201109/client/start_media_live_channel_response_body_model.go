// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMediaLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartMediaLiveChannelResponseBody
	GetRequestId() *string
}

type StartMediaLiveChannelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartMediaLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartMediaLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *StartMediaLiveChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartMediaLiveChannelResponseBody) SetRequestId(v string) *StartMediaLiveChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMediaLiveChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
