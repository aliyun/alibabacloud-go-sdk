// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyFileResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *ModifyFileResponseBody
	GetResult() map[string]interface{}
}

type ModifyFileResponseBody struct {
	// id of request
	//
	// example:
	//
	// 89B968E6-1E41-58DF-BB25-5F98ECC759CE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the index
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyFileResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *ModifyFileResponseBody) SetRequestId(v string) *ModifyFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFileResponseBody) SetResult(v map[string]interface{}) *ModifyFileResponseBody {
	s.Result = v
	return s
}

func (s *ModifyFileResponseBody) Validate() error {
	return dara.Validate(s)
}
