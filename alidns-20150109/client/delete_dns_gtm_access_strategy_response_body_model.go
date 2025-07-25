// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDnsGtmAccessStrategyResponseBody
	GetRequestId() *string
}

type DeleteDnsGtmAccessStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDnsGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *DeleteDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDnsGtmAccessStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
