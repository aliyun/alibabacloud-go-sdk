// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationGroup(v *UpdateApplicationGroupResponseBodyApplicationGroup) *UpdateApplicationGroupResponseBody
	GetApplicationGroup() *UpdateApplicationGroupResponseBodyApplicationGroup
	SetRequestId(v string) *UpdateApplicationGroupResponseBody
	GetRequestId() *string
}

type UpdateApplicationGroupResponseBody struct {
	// The information about the application group.
	ApplicationGroup *UpdateApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AA9FA778-AE4B-55EC-81CC-C46BAF08A166
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponseBody) GetApplicationGroup() *UpdateApplicationGroupResponseBodyApplicationGroup {
	return s.ApplicationGroup
}

func (s *UpdateApplicationGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationGroupResponseBody) SetApplicationGroup(v *UpdateApplicationGroupResponseBodyApplicationGroup) *UpdateApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

func (s *UpdateApplicationGroupResponseBody) SetRequestId(v string) *UpdateApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationGroupResponseBodyApplicationGroup struct {
	// The application name.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The time when the application group was created.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The ID of the region in which the related resources reside.
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
	// UpdateMyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the application group was updated.
	//
	// example:
	//
	// 2021-09-08T03:01:53Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateApplicationGroupResponseBodyApplicationGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetDeployRegionId() *string {
	return s.DeployRegionId
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetImportTagKey() *string {
	return s.ImportTagKey
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetImportTagValue() *string {
	return s.ImportTagValue
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetCreatedDate(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.CreatedDate = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetImportTagKey(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagKey = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetImportTagValue(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagValue = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetName(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetUpdatedDate(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) Validate() error {
	return dara.Validate(s)
}
