// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPreprocessJobsConsoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *SubmitPreprocessJobsConsoleRequest
	GetOwnerId() *int64
	SetPreprocessType(v string) *SubmitPreprocessJobsConsoleRequest
	GetPreprocessType() *string
	SetResourceOwnerAccount(v string) *SubmitPreprocessJobsConsoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitPreprocessJobsConsoleRequest
	GetResourceOwnerId() *int64
	SetVideoId(v string) *SubmitPreprocessJobsConsoleRequest
	GetVideoId() *string
}

type SubmitPreprocessJobsConsoleRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PreprocessType       *string `json:"PreprocessType,omitempty" xml:"PreprocessType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitPreprocessJobsConsoleRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitPreprocessJobsConsoleRequest) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsConsoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitPreprocessJobsConsoleRequest) GetPreprocessType() *string {
	return s.PreprocessType
}

func (s *SubmitPreprocessJobsConsoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitPreprocessJobsConsoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitPreprocessJobsConsoleRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitPreprocessJobsConsoleRequest) SetOwnerId(v int64) *SubmitPreprocessJobsConsoleRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleRequest) SetPreprocessType(v string) *SubmitPreprocessJobsConsoleRequest {
	s.PreprocessType = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleRequest) SetResourceOwnerAccount(v string) *SubmitPreprocessJobsConsoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleRequest) SetResourceOwnerId(v int64) *SubmitPreprocessJobsConsoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleRequest) SetVideoId(v string) *SubmitPreprocessJobsConsoleRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitPreprocessJobsConsoleRequest) Validate() error {
	return dara.Validate(s)
}
