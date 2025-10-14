// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludePaths(v []*string) *CreateInstanceSnapshotRequest
	GetExcludePaths() []*string
	SetImageUrl(v string) *CreateInstanceSnapshotRequest
	GetImageUrl() *string
	SetLabels(v []*CreateInstanceSnapshotRequestLabels) *CreateInstanceSnapshotRequest
	GetLabels() []*CreateInstanceSnapshotRequestLabels
	SetOverwrite(v bool) *CreateInstanceSnapshotRequest
	GetOverwrite() *bool
	SetSnapshotDescription(v string) *CreateInstanceSnapshotRequest
	GetSnapshotDescription() *string
	SetSnapshotName(v string) *CreateInstanceSnapshotRequest
	GetSnapshotName() *string
}

type CreateInstanceSnapshotRequest struct {
	ExcludePaths []*string `json:"ExcludePaths,omitempty" xml:"ExcludePaths,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai_product/tensorflow:py36_cpu_tf1.12_ubuntu
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// {\\"foo\\": \\"bar\\"}
	Labels    []*CreateInstanceSnapshotRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Overwrite *bool                                  `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// example:
	//
	// training_data_env
	SnapshotDescription *string `json:"SnapshotDescription,omitempty" xml:"SnapshotDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// training_data_env
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateInstanceSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotRequest) GetExcludePaths() []*string {
	return s.ExcludePaths
}

func (s *CreateInstanceSnapshotRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateInstanceSnapshotRequest) GetLabels() []*CreateInstanceSnapshotRequestLabels {
	return s.Labels
}

func (s *CreateInstanceSnapshotRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *CreateInstanceSnapshotRequest) GetSnapshotDescription() *string {
	return s.SnapshotDescription
}

func (s *CreateInstanceSnapshotRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *CreateInstanceSnapshotRequest) SetExcludePaths(v []*string) *CreateInstanceSnapshotRequest {
	s.ExcludePaths = v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetImageUrl(v string) *CreateInstanceSnapshotRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetLabels(v []*CreateInstanceSnapshotRequestLabels) *CreateInstanceSnapshotRequest {
	s.Labels = v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetOverwrite(v bool) *CreateInstanceSnapshotRequest {
	s.Overwrite = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetSnapshotDescription(v string) *CreateInstanceSnapshotRequest {
	s.SnapshotDescription = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) SetSnapshotName(v string) *CreateInstanceSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateInstanceSnapshotRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceSnapshotRequestLabels struct {
	// example:
	//
	// stsTokenOwner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 123xxxxxxxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceSnapshotRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceSnapshotRequestLabels) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotRequestLabels) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceSnapshotRequestLabels) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceSnapshotRequestLabels) SetKey(v string) *CreateInstanceSnapshotRequestLabels {
	s.Key = &v
	return s
}

func (s *CreateInstanceSnapshotRequestLabels) SetValue(v string) *CreateInstanceSnapshotRequestLabels {
	s.Value = &v
	return s
}

func (s *CreateInstanceSnapshotRequestLabels) Validate() error {
	return dara.Validate(s)
}
