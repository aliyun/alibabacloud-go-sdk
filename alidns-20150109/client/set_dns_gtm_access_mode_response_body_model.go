// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDnsGtmAccessModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDnsGtmAccessModeResponseBody
	GetRequestId() *string
}

type SetDnsGtmAccessModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDnsGtmAccessModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDnsGtmAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *SetDnsGtmAccessModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDnsGtmAccessModeResponseBody) SetRequestId(v string) *SetDnsGtmAccessModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDnsGtmAccessModeResponseBody) Validate() error {
	return dara.Validate(s)
}
