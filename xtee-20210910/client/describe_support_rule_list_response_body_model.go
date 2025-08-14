// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSupportRuleListResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DescribeSupportRuleListResponseBody
	GetResultObject() *bool
}

type DescribeSupportRuleListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribeSupportRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportRuleListResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribeSupportRuleListResponseBody) SetRequestId(v string) *DescribeSupportRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportRuleListResponseBody) SetResultObject(v bool) *DescribeSupportRuleListResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribeSupportRuleListResponseBody) Validate() error {
	return dara.Validate(s)
}
