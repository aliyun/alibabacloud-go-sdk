// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvanceConfigFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAdvanceConfigFileResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyAdvanceConfigFileResponseBody
	GetResult() map[string]interface{}
}

type ModifyAdvanceConfigFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 93A9E542-8CF8-5BA6-99AB-94C0FE520429
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyAdvanceConfigFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvanceConfigFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAdvanceConfigFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAdvanceConfigFileResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyAdvanceConfigFileResponseBody) SetRequestId(v string) *ModifyAdvanceConfigFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAdvanceConfigFileResponseBody) SetResult(v map[string]interface{}) *ModifyAdvanceConfigFileResponseBody {
	s.Result = v
	return s
}

func (s *ModifyAdvanceConfigFileResponseBody) Validate() error {
	return dara.Validate(s)
}
