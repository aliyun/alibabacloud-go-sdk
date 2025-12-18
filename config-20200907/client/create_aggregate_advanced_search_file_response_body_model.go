// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateAdvancedSearchFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAggregateAdvancedSearchFileResponseBody
	GetRequestId() *string
}

type CreateAggregateAdvancedSearchFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5F290373-2BE6-534B-8724-A33F1116958B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregateAdvancedSearchFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateAdvancedSearchFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateAdvancedSearchFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAggregateAdvancedSearchFileResponseBody) SetRequestId(v string) *CreateAggregateAdvancedSearchFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAggregateAdvancedSearchFileResponseBody) Validate() error {
	return dara.Validate(s)
}
