// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreIncrDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRestoreIncrDetailRequest
	GetClusterId() *string
	SetRestoreRecordId(v string) *DescribeRestoreIncrDetailRequest
	GetRestoreRecordId() *string
}

type DescribeRestoreIncrDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eyf188hw481xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020110514xxxx
	RestoreRecordId *string `json:"RestoreRecordId,omitempty" xml:"RestoreRecordId,omitempty"`
}

func (s DescribeRestoreIncrDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreIncrDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRestoreIncrDetailRequest) GetRestoreRecordId() *string {
	return s.RestoreRecordId
}

func (s *DescribeRestoreIncrDetailRequest) SetClusterId(v string) *DescribeRestoreIncrDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreIncrDetailRequest) SetRestoreRecordId(v string) *DescribeRestoreIncrDetailRequest {
	s.RestoreRecordId = &v
	return s
}

func (s *DescribeRestoreIncrDetailRequest) Validate() error {
	return dara.Validate(s)
}
