// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGtmAccessStrategyResponseBody
	GetRequestId() *string
}

type UpdateGtmAccessStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGtmAccessStrategyResponseBody) SetRequestId(v string) *UpdateGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGtmAccessStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
