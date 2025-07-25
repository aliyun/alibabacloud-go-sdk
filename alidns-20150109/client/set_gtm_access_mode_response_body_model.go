// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGtmAccessModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetGtmAccessModeResponseBody
	GetRequestId() *string
}

type SetGtmAccessModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGtmAccessModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetGtmAccessModeResponseBody) GoString() string {
	return s.String()
}

func (s *SetGtmAccessModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetGtmAccessModeResponseBody) SetRequestId(v string) *SetGtmAccessModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetGtmAccessModeResponseBody) Validate() error {
	return dara.Validate(s)
}
