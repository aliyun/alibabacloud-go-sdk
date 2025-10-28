// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoryDeployVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHistoryDeployVersionResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHistoryDeployVersionResponseBody
	GetMessage() *string
	SetPackageVersionList(v *ListHistoryDeployVersionResponseBodyPackageVersionList) *ListHistoryDeployVersionResponseBody
	GetPackageVersionList() *ListHistoryDeployVersionResponseBodyPackageVersionList
	SetRequestId(v string) *ListHistoryDeployVersionResponseBody
	GetRequestId() *string
}

type ListHistoryDeployVersionResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about historical deployment packages.
	PackageVersionList *ListHistoryDeployVersionResponseBodyPackageVersionList `json:"PackageVersionList,omitempty" xml:"PackageVersionList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoryDeployVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHistoryDeployVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHistoryDeployVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHistoryDeployVersionResponseBody) GetPackageVersionList() *ListHistoryDeployVersionResponseBodyPackageVersionList {
	return s.PackageVersionList
}

func (s *ListHistoryDeployVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHistoryDeployVersionResponseBody) SetCode(v int32) *ListHistoryDeployVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBody) SetMessage(v string) *ListHistoryDeployVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBody) SetPackageVersionList(v *ListHistoryDeployVersionResponseBodyPackageVersionList) *ListHistoryDeployVersionResponseBody {
	s.PackageVersionList = v
	return s
}

func (s *ListHistoryDeployVersionResponseBody) SetRequestId(v string) *ListHistoryDeployVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBody) Validate() error {
	if s.PackageVersionList != nil {
		if err := s.PackageVersionList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHistoryDeployVersionResponseBodyPackageVersionList struct {
	PackageVersion []*ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty" type:"Repeated"`
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionList) String() string {
	return dara.Prettify(s)
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionList) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionList) GetPackageVersion() []*ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	return s.PackageVersion
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionList) SetPackageVersion(v []*ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) *ListHistoryDeployVersionResponseBodyPackageVersionList {
	s.PackageVersion = v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionList) Validate() error {
	if s.PackageVersion != nil {
		for _, item := range s.PackageVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion struct {
	// The ID of the application.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the deployment package was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573627440892
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// deploy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique ID of the deployment package.
	//
	// example:
	//
	// 441beb18-da42-44dc-****-************
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The version of the application that was released by using the deployment package. This version can be used to call the RollbackApplication operation.
	//
	// example:
	//
	// 1.0
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// The URL of the deployment package.
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	// The deployment mode of the application. Valid values:
	//
	// 	- url: The application is deployed by using a JAR or WAR package.
	//
	// 	- image: The application is deployed by using an image.
	//
	// example:
	//
	// url
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the deployment package was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1573627440892
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The URL of the deployment package.
	WarUrl *string `json:"WarUrl,omitempty" xml:"WarUrl,omitempty"`
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) String() string {
	return dara.Prettify(s)
}

func (s ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetAppId() *string {
	return s.AppId
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetDescription() *string {
	return s.Description
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetId() *string {
	return s.Id
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetType() *string {
	return s.Type
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) GetWarUrl() *string {
	return s.WarUrl
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetAppId(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.AppId = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetCreateTime(v int64) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.CreateTime = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetDescription(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.Description = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetId(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.Id = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetPackageVersion(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.PackageVersion = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetPublicUrl(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.PublicUrl = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetType(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.Type = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetUpdateTime(v int64) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.UpdateTime = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) SetWarUrl(v string) *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion {
	s.WarUrl = &v
	return s
}

func (s *ListHistoryDeployVersionResponseBodyPackageVersionListPackageVersion) Validate() error {
	return dara.Validate(s)
}
