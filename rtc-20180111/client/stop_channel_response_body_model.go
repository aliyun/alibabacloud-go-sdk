// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopChannelResponseBody
	GetRequestId() *string
}

type StopChannelResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopChannelResponseBody) GoString() string {
	return s.String()
}

func (s *StopChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopChannelResponseBody) SetRequestId(v string) *StopChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
