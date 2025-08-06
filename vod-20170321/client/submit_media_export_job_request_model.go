// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobName(v string) *SubmitMediaExportJobRequest
	GetJobName() *string
	SetMatch(v string) *SubmitMediaExportJobRequest
	GetMatch() *string
	SetMediaExportConfig(v string) *SubmitMediaExportJobRequest
	GetMediaExportConfig() *string
	SetMediaType(v string) *SubmitMediaExportJobRequest
	GetMediaType() *string
	SetOwnerId(v int64) *SubmitMediaExportJobRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SubmitMediaExportJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitMediaExportJobRequest
	GetResourceOwnerId() *int64
	SetSortBy(v string) *SubmitMediaExportJobRequest
	GetSortBy() *string
}

type SubmitMediaExportJobRequest struct {
	JobName           *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Match             *string `json:"Match,omitempty" xml:"Match,omitempty"`
	MediaExportConfig *string `json:"MediaExportConfig,omitempty" xml:"MediaExportConfig,omitempty"`
	// This parameter is required.
	MediaType            *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SortBy               *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s SubmitMediaExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaExportJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaExportJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *SubmitMediaExportJobRequest) GetMatch() *string {
	return s.Match
}

func (s *SubmitMediaExportJobRequest) GetMediaExportConfig() *string {
	return s.MediaExportConfig
}

func (s *SubmitMediaExportJobRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *SubmitMediaExportJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitMediaExportJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitMediaExportJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitMediaExportJobRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SubmitMediaExportJobRequest) SetJobName(v string) *SubmitMediaExportJobRequest {
	s.JobName = &v
	return s
}

func (s *SubmitMediaExportJobRequest) SetMatch(v string) *SubmitMediaExportJobRequest {
	s.Match = &v
	return s
}

func (s *SubmitMediaExportJobRequest) SetMediaExportConfig(v string) *SubmitMediaExportJobRequest {
	s.MediaExportConfig = &v
	return s
}

func (s *SubmitMediaExportJobRequest) SetMediaType(v string) *SubmitMediaExportJobRequest {
	s.MediaType = &v
	return s
}

func (s *SubmitMediaExportJobRequest) SetOwnerId(v int64) *SubmitMediaExportJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitMediaExportJobRequest) SetResourceOwnerAccount(v string) *SubmitMediaExportJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitMediaExportJobRequest) SetResourceOwnerId(v int64) *SubmitMediaExportJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitMediaExportJobRequest) SetSortBy(v string) *SubmitMediaExportJobRequest {
	s.SortBy = &v
	return s
}

func (s *SubmitMediaExportJobRequest) Validate() error {
	return dara.Validate(s)
}
