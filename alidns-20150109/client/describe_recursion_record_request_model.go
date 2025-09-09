// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecursionRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *DescribeRecursionRecordRequest
	GetRecordId() *string
}

type DescribeRecursionRecordRequest struct {
	// example:
	//
	// 1917628665627259904
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DescribeRecursionRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecursionRecordRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeRecursionRecordRequest) SetRecordId(v string) *DescribeRecursionRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeRecursionRecordRequest) Validate() error {
	return dara.Validate(s)
}
