// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportConfigJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateExportConfigJobRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateExportConfigJobRequest
	GetRegionId() *string
}

type CreateExportConfigJobRequest struct {
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the Bastionhost instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance for which you want to export the configuration backup.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateExportConfigJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExportConfigJobRequest) GoString() string {
	return s.String()
}

func (s *CreateExportConfigJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateExportConfigJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExportConfigJobRequest) SetInstanceId(v string) *CreateExportConfigJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExportConfigJobRequest) SetRegionId(v string) *CreateExportConfigJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExportConfigJobRequest) Validate() error {
	return dara.Validate(s)
}
