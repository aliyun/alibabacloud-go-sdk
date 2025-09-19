// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHttpStatusCode(v int32) *ModifyStrategyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyStrategyResponseBody
	GetRequestId() *string
	SetResult(v *ModifyStrategyResponseBodyResult) *ModifyStrategyResponseBody
	GetResult() *ModifyStrategyResponseBodyResult
	SetSuccess(v bool) *ModifyStrategyResponseBody
	GetSuccess() *bool
}

type ModifyStrategyResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8EFA2BD9-00CD-5D69-B6B0-4EE83EAF072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned.
	Result *ModifyStrategyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStrategyResponseBody) GetResult() *ModifyStrategyResponseBodyResult {
	return s.Result
}

func (s *ModifyStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyStrategyResponseBody) SetHttpStatusCode(v int32) *ModifyStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyStrategyResponseBody) SetRequestId(v string) *ModifyStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStrategyResponseBody) SetResult(v *ModifyStrategyResponseBodyResult) *ModifyStrategyResponseBody {
	s.Result = v
	return s
}

func (s *ModifyStrategyResponseBody) SetSuccess(v bool) *ModifyStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyStrategyResponseBodyResult struct {
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 8164239
	StrategyId *int32 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ModifyStrategyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyStrategyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyStrategyResponseBodyResult) GetStrategyId() *int32 {
	return s.StrategyId
}

func (s *ModifyStrategyResponseBodyResult) SetStrategyId(v int32) *ModifyStrategyResponseBodyResult {
	s.StrategyId = &v
	return s
}

func (s *ModifyStrategyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
