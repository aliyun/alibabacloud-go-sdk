// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDeployGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *InsertDeployGroupRequest
	GetAppId() *string
	SetGroupName(v string) *InsertDeployGroupRequest
	GetGroupName() *string
	SetInitPackageVersionId(v string) *InsertDeployGroupRequest
	GetInitPackageVersionId() *string
}

type InsertDeployGroupRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413-***********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the instance group. The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The version of the initial deployment package associated with the instance group. You can call the ListHistoryDeployVersion operation to query the version. For more information, see [ListHistoryDeployVersion](https://help.aliyun.com/document_detail/149392.html).
	//
	// example:
	//
	// 441beb18-da42-44dc-****-****
	InitPackageVersionId *string `json:"InitPackageVersionId,omitempty" xml:"InitPackageVersionId,omitempty"`
}

func (s InsertDeployGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *InsertDeployGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *InsertDeployGroupRequest) GetInitPackageVersionId() *string {
	return s.InitPackageVersionId
}

func (s *InsertDeployGroupRequest) SetAppId(v string) *InsertDeployGroupRequest {
	s.AppId = &v
	return s
}

func (s *InsertDeployGroupRequest) SetGroupName(v string) *InsertDeployGroupRequest {
	s.GroupName = &v
	return s
}

func (s *InsertDeployGroupRequest) SetInitPackageVersionId(v string) *InsertDeployGroupRequest {
	s.InitPackageVersionId = &v
	return s
}

func (s *InsertDeployGroupRequest) Validate() error {
	return dara.Validate(s)
}
