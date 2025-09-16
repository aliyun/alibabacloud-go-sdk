// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexOnlineStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetIndexOnlineStrategyResponseBody
	GetRequestId() *string
	SetResult(v *GetIndexOnlineStrategyResponseBodyResult) *GetIndexOnlineStrategyResponseBody
	GetResult() *GetIndexOnlineStrategyResponseBodyResult
}

type GetIndexOnlineStrategyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE03180A-0E29-5474-8A86-33F0683294A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result *GetIndexOnlineStrategyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexOnlineStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIndexOnlineStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexOnlineStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIndexOnlineStrategyResponseBody) GetResult() *GetIndexOnlineStrategyResponseBodyResult {
	return s.Result
}

func (s *GetIndexOnlineStrategyResponseBody) SetRequestId(v string) *GetIndexOnlineStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexOnlineStrategyResponseBody) SetResult(v *GetIndexOnlineStrategyResponseBodyResult) *GetIndexOnlineStrategyResponseBody {
	s.Result = v
	return s
}

func (s *GetIndexOnlineStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIndexOnlineStrategyResponseBodyResult struct {
	// The index change rate.
	//
	// example:
	//
	// 20
	ChangeRate *int32 `json:"changeRate,omitempty" xml:"changeRate,omitempty"`
}

func (s GetIndexOnlineStrategyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetIndexOnlineStrategyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexOnlineStrategyResponseBodyResult) GetChangeRate() *int32 {
	return s.ChangeRate
}

func (s *GetIndexOnlineStrategyResponseBodyResult) SetChangeRate(v int32) *GetIndexOnlineStrategyResponseBodyResult {
	s.ChangeRate = &v
	return s
}

func (s *GetIndexOnlineStrategyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
