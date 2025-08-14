// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMediaLiveChannelResponseBody
	GetRequestId() *string
}

type DeleteMediaLiveChannelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaLiveChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaLiveChannelResponseBody) SetRequestId(v string) *DeleteMediaLiveChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaLiveChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
