// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpDBDeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelType(v string) *SubmitFpDBDeleteJobRequest
	GetDelType() *string
	SetFpDBId(v string) *SubmitFpDBDeleteJobRequest
	GetFpDBId() *string
	SetOwnerAccount(v string) *SubmitFpDBDeleteJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitFpDBDeleteJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitFpDBDeleteJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *SubmitFpDBDeleteJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitFpDBDeleteJobRequest
	GetResourceOwnerId() *int64
	SetUserData(v string) *SubmitFpDBDeleteJobRequest
	GetUserData() *string
}

type SubmitFpDBDeleteJobRequest struct {
	// The operation type. Valid values:
	//
	// 	- **Purge**: clears the media fingerprint library. The content in the library is deleted, but the library is not deleted.
	//
	// 	- **Delete**: deletes the media fingerprint library. Both the library and its content are deleted.
	//
	// 	- Default value: **Purge**.
	//
	// example:
	//
	// Purge
	DelType *string `json:"DelType,omitempty" xml:"DelType,omitempty"`
	// The ID of the media fingerprint library. You can obtain the library ID from the response parameters of the [CreateFpShotDB](https://help.aliyun.com/document_detail/170149.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	FpDBId       *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue. This ID can be used to associate the job with a notification method. To view the MPS queue ID, log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane.
	//
	// example:
	//
	// fb712a6890464059b1b2ea7c8647****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The custom data. The value can contain letters and digits and can be up to 128 bytes in length.
	//
	// example:
	//
	// example data
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitFpDBDeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpDBDeleteJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitFpDBDeleteJobRequest) GetDelType() *string {
	return s.DelType
}

func (s *SubmitFpDBDeleteJobRequest) GetFpDBId() *string {
	return s.FpDBId
}

func (s *SubmitFpDBDeleteJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitFpDBDeleteJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitFpDBDeleteJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitFpDBDeleteJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitFpDBDeleteJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitFpDBDeleteJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitFpDBDeleteJobRequest) SetDelType(v string) *SubmitFpDBDeleteJobRequest {
	s.DelType = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) SetFpDBId(v string) *SubmitFpDBDeleteJobRequest {
	s.FpDBId = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) SetOwnerAccount(v string) *SubmitFpDBDeleteJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) SetOwnerId(v int64) *SubmitFpDBDeleteJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) SetPipelineId(v string) *SubmitFpDBDeleteJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) SetResourceOwnerAccount(v string) *SubmitFpDBDeleteJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) SetResourceOwnerId(v int64) *SubmitFpDBDeleteJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) SetUserData(v string) *SubmitFpDBDeleteJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitFpDBDeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
