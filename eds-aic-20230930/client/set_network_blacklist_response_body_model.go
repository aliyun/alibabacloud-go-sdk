// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetNetworkBlacklistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetNetworkBlacklistResponseBody
	GetRequestId() *string
}

type SetNetworkBlacklistResponseBody struct {
	// example:
	//
	// 2B9E6946-0E2A-5D2B-B275-361DF81F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetNetworkBlacklistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetNetworkBlacklistResponseBody) GoString() string {
	return s.String()
}

func (s *SetNetworkBlacklistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetNetworkBlacklistResponseBody) SetRequestId(v string) *SetNetworkBlacklistResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetNetworkBlacklistResponseBody) Validate() error {
	return dara.Validate(s)
}
