// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateDatasetRequest
	GetAccessibility() *string
	SetAccessibleRoleIdList(v []*string) *UpdateDatasetRequest
	GetAccessibleRoleIdList() []*string
	SetDescription(v string) *UpdateDatasetRequest
	GetDescription() *string
	SetEdition(v string) *UpdateDatasetRequest
	GetEdition() *string
	SetMountAccessReadWriteRoleIdList(v []*string) *UpdateDatasetRequest
	GetMountAccessReadWriteRoleIdList() []*string
	SetName(v string) *UpdateDatasetRequest
	GetName() *string
	SetOptions(v string) *UpdateDatasetRequest
	GetOptions() *string
	SetSharingConfig(v *UpdateDatasetRequestSharingConfig) *UpdateDatasetRequest
	GetSharingConfig() *UpdateDatasetRequestSharingConfig
}

type UpdateDatasetRequest struct {
	// The visibility of the dataset in the workspace. Valid values:
	//
	// - `PRIVATE` (default): The dataset is visible only to its owner and administrators.
	//
	// - `PUBLIC`: The dataset is visible to all users in the workspace.
	//
	// - `ROLE_PUBLIC`: The dataset is visible to users in specific workspace roles. You must specify the roles in the `AccessibleRoleIdList` parameter. The dataset owner and administrators can always view the dataset.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// This parameter takes effect only when `Accessibility` is set to `ROLE_PUBLIC`. It specifies the list of workspace roles that can view the dataset. Role IDs that start with `PAI` are basic role IDs, and role IDs that start with `role-` are custom role IDs.
	AccessibleRoleIdList []*string `json:"AccessibleRoleIdList,omitempty" xml:"AccessibleRoleIdList,omitempty" type:"Repeated"`
	// The description of the dataset.
	//
	// example:
	//
	// This is a description of the dataset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The dataset edition. You can upgrade a dataset from `BASIC` to `ADVANCED`.
	//
	// example:
	//
	// ADVANCED
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// A list of workspace roles that have read and write permissions on the mounted dataset. Role IDs that start with `PAI` are basic role IDs, and role IDs that start with `role-` are custom role IDs. If the list contains an asterisk (`*`), all roles are granted read and write permissions.
	//
	// - To specify roles: ["PAI.AlgoOperator", "role-hiuwpd01ncrokkgp21"]
	//
	// - To specify all roles: ["\\*"]
	//
	// - To specify only the dataset creator: []
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The dataset name. For information about how to obtain the dataset name, see [ListDatasets](https://help.aliyun.com/document_detail/457222.html).
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An extended field in a JSON string format. When you use the dataset with Data Lake Compute (DLC), you can configure the `mountPath` field to specify the default mount path.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The sharing configuration of the dataset.
	//
	// if can be null:
	// true
	SharingConfig *UpdateDatasetRequestSharingConfig `json:"SharingConfig,omitempty" xml:"SharingConfig,omitempty" type:"Struct"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateDatasetRequest) GetAccessibleRoleIdList() []*string {
	return s.AccessibleRoleIdList
}

func (s *UpdateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetRequest) GetEdition() *string {
	return s.Edition
}

func (s *UpdateDatasetRequest) GetMountAccessReadWriteRoleIdList() []*string {
	return s.MountAccessReadWriteRoleIdList
}

func (s *UpdateDatasetRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateDatasetRequest) GetSharingConfig() *UpdateDatasetRequestSharingConfig {
	return s.SharingConfig
}

func (s *UpdateDatasetRequest) SetAccessibility(v string) *UpdateDatasetRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateDatasetRequest) SetAccessibleRoleIdList(v []*string) *UpdateDatasetRequest {
	s.AccessibleRoleIdList = v
	return s
}

func (s *UpdateDatasetRequest) SetDescription(v string) *UpdateDatasetRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequest) SetEdition(v string) *UpdateDatasetRequest {
	s.Edition = &v
	return s
}

func (s *UpdateDatasetRequest) SetMountAccessReadWriteRoleIdList(v []*string) *UpdateDatasetRequest {
	s.MountAccessReadWriteRoleIdList = v
	return s
}

func (s *UpdateDatasetRequest) SetName(v string) *UpdateDatasetRequest {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequest) SetOptions(v string) *UpdateDatasetRequest {
	s.Options = &v
	return s
}

func (s *UpdateDatasetRequest) SetSharingConfig(v *UpdateDatasetRequestSharingConfig) *UpdateDatasetRequest {
	s.SharingConfig = v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	if s.SharingConfig != nil {
		if err := s.SharingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDatasetRequestSharingConfig struct {
	// The sharing relationships of the dataset.
	SharedTo []*DatasetShareRelationship `json:"SharedTo,omitempty" xml:"SharedTo,omitempty" type:"Repeated"`
}

func (s UpdateDatasetRequestSharingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequestSharingConfig) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequestSharingConfig) GetSharedTo() []*DatasetShareRelationship {
	return s.SharedTo
}

func (s *UpdateDatasetRequestSharingConfig) SetSharedTo(v []*DatasetShareRelationship) *UpdateDatasetRequestSharingConfig {
	s.SharedTo = v
	return s
}

func (s *UpdateDatasetRequestSharingConfig) Validate() error {
	if s.SharedTo != nil {
		for _, item := range s.SharedTo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
