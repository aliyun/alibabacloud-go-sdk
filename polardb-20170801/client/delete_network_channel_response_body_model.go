// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkChannelResponseBody
	GetRequestId() *string
}

type DeleteNetworkChannelResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3AA69096-757C-4647-B36C-29EBC2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkChannelResponseBody) SetRequestId(v string) *DeleteNetworkChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
