// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSnapshotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v string) *SubmitSnapshotJobRequest
	GetInput() *string
	SetOwnerAccount(v string) *SubmitSnapshotJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitSnapshotJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitSnapshotJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *SubmitSnapshotJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitSnapshotJobRequest
	GetResourceOwnerId() *int64
	SetSnapshotConfig(v string) *SubmitSnapshotJobRequest
	GetSnapshotConfig() *string
	SetUserData(v string) *SubmitSnapshotJobRequest
	GetUserData() *string
}

type SubmitSnapshotJobRequest struct {
	// The information about the job input. The value must be a JSON object. You must add the Object Storage Service (OSS) bucket that stores the OSS object to be used as the job input as a media bucket in the MPS console. To add an OSS bucket as a media bucket, you can log on to the MPS console, choose Workflows > Media Buckets in the left-side navigation pane, and then click Add Bucket. After the OSS bucket is added as a media bucket, you must perform URL encoding for the OSS object. Example: `{"Bucket":"example-bucket","Location":"example-location","Object":"example%2Ftest.flv"}`. This example indicates the `"example-bucket.example-location.aliyuncs.com/example/test.flv"` object.
	//
	// > The OSS bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"example-location","Object":"example%2Ftest.flv"}
	Input        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue to which you want to submit the snapshot job. To obtain the ID, you can log on to the **MPS console*	- and choose **Global Settings*	- > **Pipelines*	- in the left-side navigation pane.
	//
	// > Make sure that an available Message Service (MNS) topic is bound to the specified MPS queue. Otherwise, the relevant messages may fail to be sent as expected.
	//
	// example:
	//
	// dd3dae411e704030b921e52698e5****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot configurations. For more information, see the "AliyunSnapshotConfig" section of the [Data types](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// > If you set the Interval parameter that is nested under SnapshotConfig, snapshots are captured at the specified intervals. The default value of the Interval parameter is 10, in seconds. If an input video is short but you specify large values for both the Num and Interval parameters, the actual number of snapshots captured may be smaller than the specified number. For example, if you set the Num parameter to 5 and the Interval parameter to 3 for a video of 10 seconds, the number of snapshots captured cannot reach 5.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"OutputFile":{"Bucket":"example-001","Location":"example-location","Object":"{Count}.jpg"},"Time":"5","Num":"10","Interval":"20"}
	SnapshotConfig *string `json:"SnapshotConfig,omitempty" xml:"SnapshotConfig,omitempty"`
	// The custom data. The custom data can contain letters, digits, and hyphens (-) and be up to 1,024 bytes in size. The custom data cannot start with a special character.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSnapshotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitSnapshotJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitSnapshotJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitSnapshotJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitSnapshotJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitSnapshotJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitSnapshotJobRequest) GetSnapshotConfig() *string {
	return s.SnapshotConfig
}

func (s *SubmitSnapshotJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSnapshotJobRequest) SetInput(v string) *SubmitSnapshotJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetOwnerAccount(v string) *SubmitSnapshotJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetOwnerId(v int64) *SubmitSnapshotJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetPipelineId(v string) *SubmitSnapshotJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetResourceOwnerAccount(v string) *SubmitSnapshotJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetResourceOwnerId(v int64) *SubmitSnapshotJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSnapshotConfig(v string) *SubmitSnapshotJobRequest {
	s.SnapshotConfig = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetUserData(v string) *SubmitSnapshotJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSnapshotJobRequest) Validate() error {
	return dara.Validate(s)
}
