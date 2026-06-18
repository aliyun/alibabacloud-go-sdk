// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagValue(v string) *DescribeQuotaRequest
	GetTagValue() *string
}

type DescribeQuotaRequest struct {
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s DescribeQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeQuotaRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeQuotaRequest) SetTagValue(v string) *DescribeQuotaRequest {
	s.TagValue = &v
	return s
}

func (s *DescribeQuotaRequest) Validate() error {
	return dara.Validate(s)
}
