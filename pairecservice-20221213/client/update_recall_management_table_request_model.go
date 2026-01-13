// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecallManagementTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableDataSizeFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest
	GetEnableDataSizeFluctuationThreshold() *bool
	SetEnableRowCountFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest
	GetEnableRowCountFluctuationThreshold() *bool
	SetIndexVersionId(v string) *UpdateRecallManagementTableRequest
	GetIndexVersionId() *string
	SetInstanceId(v string) *UpdateRecallManagementTableRequest
	GetInstanceId() *string
	SetMaxDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMaxDataSizeFluctuationThreshold() *int32
	SetMaxRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMaxRowCountFluctuationThreshold() *int32
	SetMinDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMinDataSizeFluctuationThreshold() *int32
	SetMinRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest
	GetMinRowCountFluctuationThreshold() *int32
}

type UpdateRecallManagementTableRequest struct {
	EnableDataSizeFluctuationThreshold *bool   `json:"EnableDataSizeFluctuationThreshold,omitempty" xml:"EnableDataSizeFluctuationThreshold,omitempty"`
	EnableRowCountFluctuationThreshold *bool   `json:"EnableRowCountFluctuationThreshold,omitempty" xml:"EnableRowCountFluctuationThreshold,omitempty"`
	IndexVersionId                     *string `json:"IndexVersionId,omitempty" xml:"IndexVersionId,omitempty"`
	// This parameter is required.
	InstanceId                      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxDataSizeFluctuationThreshold *int32  `json:"MaxDataSizeFluctuationThreshold,omitempty" xml:"MaxDataSizeFluctuationThreshold,omitempty"`
	MaxRowCountFluctuationThreshold *int32  `json:"MaxRowCountFluctuationThreshold,omitempty" xml:"MaxRowCountFluctuationThreshold,omitempty"`
	MinDataSizeFluctuationThreshold *int32  `json:"MinDataSizeFluctuationThreshold,omitempty" xml:"MinDataSizeFluctuationThreshold,omitempty"`
	MinRowCountFluctuationThreshold *int32  `json:"MinRowCountFluctuationThreshold,omitempty" xml:"MinRowCountFluctuationThreshold,omitempty"`
}

func (s UpdateRecallManagementTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecallManagementTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecallManagementTableRequest) GetEnableDataSizeFluctuationThreshold() *bool {
	return s.EnableDataSizeFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetEnableRowCountFluctuationThreshold() *bool {
	return s.EnableRowCountFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetIndexVersionId() *string {
	return s.IndexVersionId
}

func (s *UpdateRecallManagementTableRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRecallManagementTableRequest) GetMaxDataSizeFluctuationThreshold() *int32 {
	return s.MaxDataSizeFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetMaxRowCountFluctuationThreshold() *int32 {
	return s.MaxRowCountFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetMinDataSizeFluctuationThreshold() *int32 {
	return s.MinDataSizeFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) GetMinRowCountFluctuationThreshold() *int32 {
	return s.MinRowCountFluctuationThreshold
}

func (s *UpdateRecallManagementTableRequest) SetEnableDataSizeFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest {
	s.EnableDataSizeFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetEnableRowCountFluctuationThreshold(v bool) *UpdateRecallManagementTableRequest {
	s.EnableRowCountFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetIndexVersionId(v string) *UpdateRecallManagementTableRequest {
	s.IndexVersionId = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetInstanceId(v string) *UpdateRecallManagementTableRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMaxDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MaxDataSizeFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMaxRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MaxRowCountFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMinDataSizeFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MinDataSizeFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) SetMinRowCountFluctuationThreshold(v int32) *UpdateRecallManagementTableRequest {
	s.MinRowCountFluctuationThreshold = &v
	return s
}

func (s *UpdateRecallManagementTableRequest) Validate() error {
	return dara.Validate(s)
}
