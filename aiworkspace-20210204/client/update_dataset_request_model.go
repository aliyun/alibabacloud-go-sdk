// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The description of the dataset.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Edition     *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The list of role names in the workspace that have read and write permissions on the mounted database. The names starting with PAI are basic role names, and the names starting with role- are custom role names. If the list contains asterisks (\\*), all roles have read and write permissions.
	//
	// 	- If you set the value to ["PAI.AlgoOperator", "role-hiuwpd01ncrokkgp21"], the account of the specified role is granted the read and write permissions.
	//
	// 	- If you set the value to ["\\*"], all accounts are granted the read and write permissions.
	//
	// 	- If you set the value to [], only the creator of the dataset has the read and write permissions.
	MountAccessReadWriteRoleIdList []*string `json:"MountAccessReadWriteRoleIdList,omitempty" xml:"MountAccessReadWriteRoleIdList,omitempty" type:"Repeated"`
	// The dataset name. You can call [ListDatasets](https://help.aliyun.com/document_detail/457222.html) to obtain the dataset name.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The extended field, which is a JSON string. When you use the dataset in Deep Learning Containers (DLC), you can set mountPath to specify the default mount path of the dataset.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
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
