// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRuleResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *DeleteRuleResponseBody
	GetResultObject() *bool
}

type DeleteRuleResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRuleResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetResultObject(v bool) *DeleteRuleResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DeleteRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
