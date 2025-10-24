// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaInfoJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *SubmitMediaInfoJobRequest
	GetAsync() *bool
	SetConfig(v string) *SubmitMediaInfoJobRequest
	GetConfig() *string
	SetInput(v string) *SubmitMediaInfoJobRequest
	GetInput() *string
	SetOwnerAccount(v string) *SubmitMediaInfoJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitMediaInfoJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitMediaInfoJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *SubmitMediaInfoJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitMediaInfoJobRequest
	GetResourceOwnerId() *int64
	SetUserData(v string) *SubmitMediaInfoJobRequest
	GetUserData() *string
}

type SubmitMediaInfoJobRequest struct {
	// Specifies whether to enable the asynchronous mode for the job. We recommend that you set this parameter to true. Valid values:
	//
	// 	- **true**: enables the asynchronous mode.
	//
	// 	- **false**: does not enable the asynchronous mode.
	//
	// example:
	//
	// true
	Async  *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The information about the input media file. The value is a JSON string. You must perform the following operations to add the OSS bucket in which the input media file is stored as a media bucket: Log on to the **MPS console**, choose **Workflows*	- > **Media Buckets*	- in the left-side navigation pane, and then click **Add Bucket**. After you add the OSS bucket as a media bucket, you must perform URL encoding for the OSS object. For example, `{"Bucket":"example-bucket","Location":"example-location","Object":"example%2Fexample.flv"}` indicates the `example-bucket.example-location.aliyuncs.com/example/example.flv` file.
	//
	// > The OSS bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Bucket":"example-bucket","Location":"example-location","Object":"example%2Fexample.flv"}
	Input        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue to which the job was submitted. For more information, see [Terms](https://help.aliyun.com/document_detail/29197.html).
	//
	// 	- To view the ID of the MPS queue, log on to the [MPS console](https://mps.console.aliyun.com/overview) and choose **Global Settings*	- > **MPS queue and Callback*	- in the left-side navigation pane. On the MPS queue and Callback page, you can view the ID of an MPS queue or create an MPS queue.
	//
	// 	- If you want to receive asynchronous message notifications, associate an MNS queue or topic with the MPS queue. For more information, see [Receive message notifications](https://www.alibabacloud.com/help/en/mps/receive-message-notifications/?spm=a2c63.p38356.0.0.b48576d2jxNSca).
	//
	// example:
	//
	// 88c6ca184c0e432bbf5b665e2a15****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The custom data. The custom data can contain letters, digits, and hyphens (-), and can be up to 1,024 bytes in length. The custom data cannot start with a special character.
	//
	// example:
	//
	// testid-001
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaInfoJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaInfoJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaInfoJobRequest) GetAsync() *bool {
	return s.Async
}

func (s *SubmitMediaInfoJobRequest) GetConfig() *string {
	return s.Config
}

func (s *SubmitMediaInfoJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitMediaInfoJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitMediaInfoJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitMediaInfoJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaInfoJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitMediaInfoJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitMediaInfoJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaInfoJobRequest) SetAsync(v bool) *SubmitMediaInfoJobRequest {
	s.Async = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetConfig(v string) *SubmitMediaInfoJobRequest {
	s.Config = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetInput(v string) *SubmitMediaInfoJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetOwnerAccount(v string) *SubmitMediaInfoJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetOwnerId(v int64) *SubmitMediaInfoJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetPipelineId(v string) *SubmitMediaInfoJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetResourceOwnerAccount(v string) *SubmitMediaInfoJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetResourceOwnerId(v int64) *SubmitMediaInfoJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) SetUserData(v string) *SubmitMediaInfoJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaInfoJobRequest) Validate() error {
	return dara.Validate(s)
}
