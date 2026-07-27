// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOUsageDetailExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeMOUsageDetailExportRequest
	GetInstanceId() *string
}

type DescribeMOUsageDetailExportRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rds_copilot**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeMOUsageDetailExportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOUsageDetailExportRequest) GoString() string {
	return s.String()
}

func (s *DescribeMOUsageDetailExportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMOUsageDetailExportRequest) SetInstanceId(v string) *DescribeMOUsageDetailExportRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMOUsageDetailExportRequest) Validate() error {
	return dara.Validate(s)
}
