// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWhiteListStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveWhiteListStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v int64) *SaveWhiteListStrategyResponseBody
	GetStrategyId() *int64
}

type SaveWhiteListStrategyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5DFD6277-CC36-57F7-ACE6-F5952XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the application whitelist policy.
	//
	// example:
	//
	// 8634
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s SaveWhiteListStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveWhiteListStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWhiteListStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveWhiteListStrategyResponseBody) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *SaveWhiteListStrategyResponseBody) SetRequestId(v string) *SaveWhiteListStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWhiteListStrategyResponseBody) SetStrategyId(v int64) *SaveWhiteListStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *SaveWhiteListStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
