// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int64) *UpdateDatasetVersionRequest
	GetDataCount() *int64
	SetDataSize(v int64) *UpdateDatasetVersionRequest
	GetDataSize() *int64
	SetDatasetTaskRamRole(v string) *UpdateDatasetVersionRequest
	GetDatasetTaskRamRole() *string
	SetDescription(v string) *UpdateDatasetVersionRequest
	GetDescription() *string
	SetOptions(v string) *UpdateDatasetVersionRequest
	GetOptions() *string
	SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *UpdateDatasetVersionRequest
	GetUserMetricsEndpoints() []*UserMetricsEndpoint
}

type UpdateDatasetVersionRequest struct {
	// The number of dataset files.
	//
	// example:
	//
	// 100
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The size of the space occupied by dataset files. Unit: bytes.
	//
	// example:
	//
	// 100000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// DatasetTaskRamRole
	//
	// example:
	//
	// acs:ram::1234567890123456:role/role-name
	DatasetTaskRamRole *string `json:"DatasetTaskRamRole,omitempty" xml:"DatasetTaskRamRole,omitempty"`
	// The custom description of the dataset, which is used to distinguish different datasets.
	//
	// example:
	//
	// This is a description of a dataset version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended field in JsonString format. When DLC uses the dataset, you can specify the default mount path of the dataset by configuring the mountPath field.
	//
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options              *string                `json:"Options,omitempty" xml:"Options,omitempty"`
	UserMetricsEndpoints []*UserMetricsEndpoint `json:"UserMetricsEndpoints,omitempty" xml:"UserMetricsEndpoints,omitempty" type:"Repeated"`
}

func (s UpdateDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetVersionRequest) GetDataCount() *int64 {
	return s.DataCount
}

func (s *UpdateDatasetVersionRequest) GetDataSize() *int64 {
	return s.DataSize
}

func (s *UpdateDatasetVersionRequest) GetDatasetTaskRamRole() *string {
	return s.DatasetTaskRamRole
}

func (s *UpdateDatasetVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetVersionRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateDatasetVersionRequest) GetUserMetricsEndpoints() []*UserMetricsEndpoint {
	return s.UserMetricsEndpoints
}

func (s *UpdateDatasetVersionRequest) SetDataCount(v int64) *UpdateDatasetVersionRequest {
	s.DataCount = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetDataSize(v int64) *UpdateDatasetVersionRequest {
	s.DataSize = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetDatasetTaskRamRole(v string) *UpdateDatasetVersionRequest {
	s.DatasetTaskRamRole = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetDescription(v string) *UpdateDatasetVersionRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetOptions(v string) *UpdateDatasetVersionRequest {
	s.Options = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetUserMetricsEndpoints(v []*UserMetricsEndpoint) *UpdateDatasetVersionRequest {
	s.UserMetricsEndpoints = v
	return s
}

func (s *UpdateDatasetVersionRequest) Validate() error {
	if s.UserMetricsEndpoints != nil {
		for _, item := range s.UserMetricsEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
