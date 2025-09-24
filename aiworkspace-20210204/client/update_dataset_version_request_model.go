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
	SetDescription(v string) *UpdateDatasetVersionRequest
	GetDescription() *string
	SetOptions(v string) *UpdateDatasetVersionRequest
	GetOptions() *string
}

type UpdateDatasetVersionRequest struct {
	// example:
	//
	// 100
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// example:
	//
	// 100000
	DataSize    *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
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

func (s *UpdateDatasetVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetVersionRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateDatasetVersionRequest) SetDataCount(v int64) *UpdateDatasetVersionRequest {
	s.DataCount = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetDataSize(v int64) *UpdateDatasetVersionRequest {
	s.DataSize = &v
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

func (s *UpdateDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
