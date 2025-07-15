// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiskEncryptionServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateDiskEncryptionServiceRequest
	GetRegionId() *string
}

type CreateDiskEncryptionServiceRequest struct {
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDiskEncryptionServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiskEncryptionServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskEncryptionServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiskEncryptionServiceRequest) SetRegionId(v string) *CreateDiskEncryptionServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiskEncryptionServiceRequest) Validate() error {
	return dara.Validate(s)
}
