// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleOutApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ScaleOutApplicationRequest
	GetAppId() *string
	SetDeployGroup(v string) *ScaleOutApplicationRequest
	GetDeployGroup() *string
	SetEcuInfo(v string) *ScaleOutApplicationRequest
	GetEcuInfo() *string
}

type ScaleOutApplicationRequest struct {
	// The ID of the application that you want to scale out. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413*****************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the instance group where the application you want to scale out is deployed. You can call the QueryApplicationStatus operation to query the group ID. For more information, see [QueryApplicationStatus](https://help.aliyun.com/document_detail/149394.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 4f1fb6f5-6769-4bd6******************
	DeployGroup *string `json:"DeployGroup,omitempty" xml:"DeployGroup,omitempty"`
	// The ID of the elastic compute unit (ECU) that corresponds to the Elastic Compute Service (ECS) instance to be added to the instance group for scale-out. You can call the ListScaleOutEcu operation to query the ECU ID. For more information, see [ListScaleOutEcu](https://help.aliyun.com/document_detail/149371.html). Separate multiple ECU IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 8123db90-880f-486f-b***************
	EcuInfo *string `json:"EcuInfo,omitempty" xml:"EcuInfo,omitempty"`
}

func (s ScaleOutApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleOutApplicationRequest) GoString() string {
	return s.String()
}

func (s *ScaleOutApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *ScaleOutApplicationRequest) GetDeployGroup() *string {
	return s.DeployGroup
}

func (s *ScaleOutApplicationRequest) GetEcuInfo() *string {
	return s.EcuInfo
}

func (s *ScaleOutApplicationRequest) SetAppId(v string) *ScaleOutApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ScaleOutApplicationRequest) SetDeployGroup(v string) *ScaleOutApplicationRequest {
	s.DeployGroup = &v
	return s
}

func (s *ScaleOutApplicationRequest) SetEcuInfo(v string) *ScaleOutApplicationRequest {
	s.EcuInfo = &v
	return s
}

func (s *ScaleOutApplicationRequest) Validate() error {
	return dara.Validate(s)
}
