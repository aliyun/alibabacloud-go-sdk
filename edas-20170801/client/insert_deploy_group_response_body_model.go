// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDeployGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertDeployGroupResponseBody
	GetCode() *int32
	SetDeployGroupEntity(v *InsertDeployGroupResponseBodyDeployGroupEntity) *InsertDeployGroupResponseBody
	GetDeployGroupEntity() *InsertDeployGroupResponseBodyDeployGroupEntity
	SetMessage(v string) *InsertDeployGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertDeployGroupResponseBody
	GetRequestId() *string
}

type InsertDeployGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the instance group.
	DeployGroupEntity *InsertDeployGroupResponseBodyDeployGroupEntity `json:"DeployGroupEntity,omitempty" xml:"DeployGroupEntity,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3RD9-D3FRE****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertDeployGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertDeployGroupResponseBody) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertDeployGroupResponseBody) GetDeployGroupEntity() *InsertDeployGroupResponseBodyDeployGroupEntity {
	return s.DeployGroupEntity
}

func (s *InsertDeployGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertDeployGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertDeployGroupResponseBody) SetCode(v int32) *InsertDeployGroupResponseBody {
	s.Code = &v
	return s
}

func (s *InsertDeployGroupResponseBody) SetDeployGroupEntity(v *InsertDeployGroupResponseBodyDeployGroupEntity) *InsertDeployGroupResponseBody {
	s.DeployGroupEntity = v
	return s
}

func (s *InsertDeployGroupResponseBody) SetMessage(v string) *InsertDeployGroupResponseBody {
	s.Message = &v
	return s
}

func (s *InsertDeployGroupResponseBody) SetRequestId(v string) *InsertDeployGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertDeployGroupResponseBody) Validate() error {
	if s.DeployGroupEntity != nil {
		if err := s.DeployGroupEntity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertDeployGroupResponseBodyDeployGroupEntity struct {
	// The ID of the application.
	//
	// example:
	//
	// 3616cdca-4f92-4413-***********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The version of the deployment package for the application.
	//
	// 	- If the application is deployed, a string of random numbers is returned.
	//
	// 	- If the application is not deployed, the return value is empty.
	//
	// example:
	//
	// ****f4c50-16ee-a02b-667*****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// 0d247b93-8d62-4e34***********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the instance group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573627695779
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The name of the instance group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the instance group. Valid values:
	//
	// 	- 0: the default group.
	//
	// 	- 1: a group for which canary traffic management is not enabled.
	//
	// 	- 2: a group for which canary traffic management is enabled.
	//
	// example:
	//
	// 1
	GroupType *int32 `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// 577f4c50-16ee-43d8-a02b-667*********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The version of the deployment package that was used to deploy an application in the instance group.
	//
	// 	- If an application is deployed in the instance group, a string of random numbers is returned.
	//
	// 	- If no application is deployed in the instance group, the return value is empty.
	//
	// example:
	//
	// ****7b93-8d62-4e34***********
	PackageVersionId *string `json:"PackageVersionId,omitempty" xml:"PackageVersionId,omitempty"`
	// The time when the instance group was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573627695779
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s InsertDeployGroupResponseBodyDeployGroupEntity) String() string {
	return dara.Prettify(s)
}

func (s InsertDeployGroupResponseBodyDeployGroupEntity) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetAppId() *string {
	return s.AppId
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetClusterId() *string {
	return s.ClusterId
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetGroupName() *string {
	return s.GroupName
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetGroupType() *int32 {
	return s.GroupType
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetId() *string {
	return s.Id
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetPackageVersionId() *string {
	return s.PackageVersionId
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetAppId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.AppId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetAppVersionId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.AppVersionId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetClusterId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.ClusterId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetCreateTime(v int64) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.CreateTime = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetGroupName(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.GroupName = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetGroupType(v int32) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.GroupType = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.Id = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetPackageVersionId(v string) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.PackageVersionId = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) SetUpdateTime(v int64) *InsertDeployGroupResponseBodyDeployGroupEntity {
	s.UpdateTime = &v
	return s
}

func (s *InsertDeployGroupResponseBodyDeployGroupEntity) Validate() error {
	return dara.Validate(s)
}
