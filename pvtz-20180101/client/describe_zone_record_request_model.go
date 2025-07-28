// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *DescribeZoneRecordRequest
	GetRecordId() *int64
}

type DescribeZoneRecordRequest struct {
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5808
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DescribeZoneRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DescribeZoneRecordRequest) SetRecordId(v int64) *DescribeZoneRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeZoneRecordRequest) Validate() error {
	return dara.Validate(s)
}
