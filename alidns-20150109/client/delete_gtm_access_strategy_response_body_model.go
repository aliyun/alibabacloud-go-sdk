// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGtmAccessStrategyResponseBody
	GetRequestId() *string
}

type DeleteGtmAccessStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGtmAccessStrategyResponseBody) SetRequestId(v string) *DeleteGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGtmAccessStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
