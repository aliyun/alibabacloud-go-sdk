// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpFileDeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileIds(v string) *SubmitFpFileDeleteJobRequest
	GetFileIds() *string
	SetFpDBId(v string) *SubmitFpFileDeleteJobRequest
	GetFpDBId() *string
	SetOwnerAccount(v string) *SubmitFpFileDeleteJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitFpFileDeleteJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitFpFileDeleteJobRequest
	GetPipelineId() *string
	SetPrimaryKeys(v string) *SubmitFpFileDeleteJobRequest
	GetPrimaryKeys() *string
	SetResourceOwnerAccount(v string) *SubmitFpFileDeleteJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitFpFileDeleteJobRequest
	GetResourceOwnerId() *int64
	SetUserData(v string) *SubmitFpFileDeleteJobRequest
	GetUserData() *string
}

type SubmitFpFileDeleteJobRequest struct {
	// The IDs of the media files that you want to delete. Separate multiple file IDs with commas (,). You can delete up to 200 media files at a time. You can obtain media file IDs from the response parameters of the [ListFpShotFiles](https://help.aliyun.com/document_detail/209266.html) operation.
	//
	// example:
	//
	// 41e6536e4f2250e2e9bf26cdea19****
	FileIds *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The ID of the media fingerprint library. You can obtain the library ID from the response parameters of the [CreateFpShotDB](https://help.aliyun.com/document_detail/170149.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88c6ca184c0e432bbf5b665e2a15****
	FpDBId       *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue to which the job is submitted. The MPS queue is bound with a notification method. To view the MPS queue ID, log on to the **MPS console*	- and choose **Global Settings*	- > **MPS queue and Callback*	- in the left-side navigation pane.
	//
	// example:
	//
	// ed450ea0bfbd41e29f80a401fb4d****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The primary keys of the files to be deleted. Separate multiple primary keys with commas (,). You can delete up to 200 primary keys at a time. You can obtain the primary keys of media files from the response parameters of the [ListFpShotFiles](https://help.aliyun.com/document_detail/209266.html) operation.
	//
	// >  This parameter is available only in the China (Beijing), China (Hangzhou), and China (Shanghai) regions.
	//
	// example:
	//
	// 24e0fba7188fae707e146esa54****
	PrimaryKeys          *string `json:"PrimaryKeys,omitempty" xml:"PrimaryKeys,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The custom data. The value can contain letters and digits and can be up to 128 bytes in length.
	//
	// example:
	//
	// example data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitFpFileDeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpFileDeleteJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitFpFileDeleteJobRequest) GetFileIds() *string {
	return s.FileIds
}

func (s *SubmitFpFileDeleteJobRequest) GetFpDBId() *string {
	return s.FpDBId
}

func (s *SubmitFpFileDeleteJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitFpFileDeleteJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitFpFileDeleteJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitFpFileDeleteJobRequest) GetPrimaryKeys() *string {
	return s.PrimaryKeys
}

func (s *SubmitFpFileDeleteJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitFpFileDeleteJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitFpFileDeleteJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitFpFileDeleteJobRequest) SetFileIds(v string) *SubmitFpFileDeleteJobRequest {
	s.FileIds = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetFpDBId(v string) *SubmitFpFileDeleteJobRequest {
	s.FpDBId = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetOwnerAccount(v string) *SubmitFpFileDeleteJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetOwnerId(v int64) *SubmitFpFileDeleteJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetPipelineId(v string) *SubmitFpFileDeleteJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetPrimaryKeys(v string) *SubmitFpFileDeleteJobRequest {
	s.PrimaryKeys = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetResourceOwnerAccount(v string) *SubmitFpFileDeleteJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetResourceOwnerId(v int64) *SubmitFpFileDeleteJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) SetUserData(v string) *SubmitFpFileDeleteJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitFpFileDeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
