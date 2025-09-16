// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexOnlineStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIndexOnlineStrategyResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyIndexOnlineStrategyResponseBody
	GetResult() map[string]interface{}
}

type ModifyIndexOnlineStrategyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyIndexOnlineStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexOnlineStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIndexOnlineStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIndexOnlineStrategyResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyIndexOnlineStrategyResponseBody) SetRequestId(v string) *ModifyIndexOnlineStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIndexOnlineStrategyResponseBody) SetResult(v map[string]interface{}) *ModifyIndexOnlineStrategyResponseBody {
	s.Result = v
	return s
}

func (s *ModifyIndexOnlineStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
