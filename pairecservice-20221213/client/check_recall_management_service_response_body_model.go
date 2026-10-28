// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRecallManagementServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckRecallManagementServiceResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *CheckRecallManagementServiceResponseBody
	GetResult() map[string]interface{}
}

type CheckRecallManagementServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result details.
	Result map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckRecallManagementServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckRecallManagementServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRecallManagementServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckRecallManagementServiceResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *CheckRecallManagementServiceResponseBody) SetRequestId(v string) *CheckRecallManagementServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRecallManagementServiceResponseBody) SetResult(v map[string]interface{}) *CheckRecallManagementServiceResponseBody {
	s.Result = v
	return s
}

func (s *CheckRecallManagementServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
