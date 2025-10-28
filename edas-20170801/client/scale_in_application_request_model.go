// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleInApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ScaleInApplicationRequest
	GetAppId() *string
	SetEccInfo(v string) *ScaleInApplicationRequest
	GetEccInfo() *string
	SetForceStatus(v bool) *ScaleInApplicationRequest
	GetForceStatus() *bool
}

type ScaleInApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-44***********************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the elastic compute container (ECC) that corresponds to the Elastic Compute Service (ECS) instance to be removed for the application. Separate multiple ECC IDs with commas (,). You can call the QueryApplicationStatus operation to query the ECC ID. For more information, see [QueryApplicationStatus](https://help.aliyun.com/document_detail/149394.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 7040f221-42df-48e8-9*******************
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
	// Specifies whether to forcibly unpublish the application from the ECS instance. You need to set this parameter to true only if the ECS instance expires. In normal cases, you do not need to set this parameter to true.
	//
	// example:
	//
	// false
	ForceStatus *bool `json:"ForceStatus,omitempty" xml:"ForceStatus,omitempty"`
}

func (s ScaleInApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleInApplicationRequest) GoString() string {
	return s.String()
}

func (s *ScaleInApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *ScaleInApplicationRequest) GetEccInfo() *string {
	return s.EccInfo
}

func (s *ScaleInApplicationRequest) GetForceStatus() *bool {
	return s.ForceStatus
}

func (s *ScaleInApplicationRequest) SetAppId(v string) *ScaleInApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ScaleInApplicationRequest) SetEccInfo(v string) *ScaleInApplicationRequest {
	s.EccInfo = &v
	return s
}

func (s *ScaleInApplicationRequest) SetForceStatus(v bool) *ScaleInApplicationRequest {
	s.ForceStatus = &v
	return s
}

func (s *ScaleInApplicationRequest) Validate() error {
	return dara.Validate(s)
}
