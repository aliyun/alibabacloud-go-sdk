// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteChannelResponseBody
	GetRequestId() *string
}

type DeleteChannelResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChannelResponseBody) SetRequestId(v string) *DeleteChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
