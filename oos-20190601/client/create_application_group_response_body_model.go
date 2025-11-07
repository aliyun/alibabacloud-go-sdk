// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationGroup(v *CreateApplicationGroupResponseBodyApplicationGroup) *CreateApplicationGroupResponseBody
	GetApplicationGroup() *CreateApplicationGroupResponseBodyApplicationGroup
	SetRequestId(v string) *CreateApplicationGroupResponseBody
	GetRequestId() *string
}

type CreateApplicationGroupResponseBody struct {
	// The information about the application group.
	ApplicationGroup *CreateApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0E6BEBD3-7F9E-5878-834B-097633AB5F33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponseBody) GetApplicationGroup() *CreateApplicationGroupResponseBodyApplicationGroup {
	return s.ApplicationGroup
}

func (s *CreateApplicationGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationGroupResponseBody) SetApplicationGroup(v *CreateApplicationGroupResponseBodyApplicationGroup) *CreateApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

func (s *CreateApplicationGroupResponseBody) SetRequestId(v string) *CreateApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationGroupResponseBody) Validate() error {
	if s.ApplicationGroup != nil {
		if err := s.ApplicationGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationGroupResponseBodyApplicationGroup struct {
	// The application name.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the application group in CloudMonitor.
	//
	// example:
	//
	// 1245768
	CmsGroupId *string `json:"CmsGroupId,omitempty" xml:"CmsGroupId,omitempty"`
	// The time when the application group was created.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the region in which the related sources reside.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the application group was updated.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateApplicationGroupResponseBodyApplicationGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetCmsGroupId() *string {
	return s.CmsGroupId
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetImportTagKey() *string {
	return s.ImportTagKey
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetImportTagValue() *string {
	return s.ImportTagValue
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetName() *string {
	return s.Name
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetCmsGroupId(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.CmsGroupId = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetCreateDate(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetImportTagKey(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagKey = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetImportTagValue(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagValue = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetName(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetUpdateDate(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.UpdateDate = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) Validate() error {
	return dara.Validate(s)
}
