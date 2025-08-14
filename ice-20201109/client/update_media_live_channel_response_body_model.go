// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMediaLiveChannelResponseBody
	GetRequestId() *string
}

type UpdateMediaLiveChannelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaLiveChannelResponseBody) SetRequestId(v string) *UpdateMediaLiveChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaLiveChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
