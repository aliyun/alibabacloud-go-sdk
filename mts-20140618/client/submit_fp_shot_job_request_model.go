// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitFpShotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFpShotConfig(v string) *SubmitFpShotJobRequest
	GetFpShotConfig() *string
	SetInput(v string) *SubmitFpShotJobRequest
	GetInput() *string
	SetOwnerAccount(v string) *SubmitFpShotJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitFpShotJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitFpShotJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *SubmitFpShotJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitFpShotJobRequest
	GetResourceOwnerId() *int64
	SetUserData(v string) *SubmitFpShotJobRequest
	GetUserData() *string
}

type SubmitFpShotJobRequest struct {
	// The configurations of the media fingerprint analysis job. The value is a JSON object. For more information, see the "FpShotConfig" section of the [Parameter details](https://help.aliyun.com/document_detail/93568.html) topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//       "PrimaryKey": "12345****",
	//
	//       "SaveType": "save",
	//
	//       "FpDBId": "417f2ada5999daf****"
	//
	// }
	FpShotConfig *string `json:"FpShotConfig,omitempty" xml:"FpShotConfig,omitempty"`
	// The OSS URL of the job input. The value is a JSON object. You can query the OSS URL in the OSS or MPS console.
	//
	// > The OSS bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket-****","Location":"oss-cn-shanghai","Object":"example-****.flv"}
	Input        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue. This ID can be used to associate the job with a notification method. To view the ID of the MPS queue, perform the following steps: Log on to the **MPS console**. In the left-side navigation pane, choose **Global Settings*	- > **Pipelines**.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The custom data. The value can be up to 128 bytes in length and cannot start with a special character.
	//
	// example:
	//
	// testid-****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitFpShotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitFpShotJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitFpShotJobRequest) GetFpShotConfig() *string {
	return s.FpShotConfig
}

func (s *SubmitFpShotJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitFpShotJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitFpShotJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitFpShotJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitFpShotJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitFpShotJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitFpShotJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitFpShotJobRequest) SetFpShotConfig(v string) *SubmitFpShotJobRequest {
	s.FpShotConfig = &v
	return s
}

func (s *SubmitFpShotJobRequest) SetInput(v string) *SubmitFpShotJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitFpShotJobRequest) SetOwnerAccount(v string) *SubmitFpShotJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitFpShotJobRequest) SetOwnerId(v int64) *SubmitFpShotJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitFpShotJobRequest) SetPipelineId(v string) *SubmitFpShotJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitFpShotJobRequest) SetResourceOwnerAccount(v string) *SubmitFpShotJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitFpShotJobRequest) SetResourceOwnerId(v int64) *SubmitFpShotJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitFpShotJobRequest) SetUserData(v string) *SubmitFpShotJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitFpShotJobRequest) Validate() error {
	return dara.Validate(s)
}
