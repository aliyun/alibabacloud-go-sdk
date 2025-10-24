// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v string) *SubmitJobsRequest
	GetInput() *string
	SetOutputBucket(v string) *SubmitJobsRequest
	GetOutputBucket() *string
	SetOutputLocation(v string) *SubmitJobsRequest
	GetOutputLocation() *string
	SetOutputs(v string) *SubmitJobsRequest
	GetOutputs() *string
	SetOwnerAccount(v string) *SubmitJobsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitJobsRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitJobsRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *SubmitJobsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitJobsRequest
	GetResourceOwnerId() *int64
}

type SubmitJobsRequest struct {
	// The information about the input file. For more information, see the "Input" section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// >
	//
	// 	- The path of an Object Storage Service (OSS) object must be URL-encoded in UTF-8 before you use the path in MPS.
	//
	// 	- The OSS bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
	//
	// example:
	//
	// a/b/c/test-cn.srt
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the OSS bucket that stores the output file.
	//
	// 	- For more information about the term bucket, see [Terms](https://help.aliyun.com/document_detail/31827.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleBucket
	OutputBucket *string `json:"OutputBucket,omitempty" xml:"OutputBucket,omitempty"`
	// The region in which the OSS bucket that stores the output file resides.
	//
	// 	- The OSS bucket must reside in the same region as MPS.
	//
	// 	- For more information about the term bucket, see [Terms](https://help.aliyun.com/document_detail/31827.html).
	//
	// example:
	//
	// oss-cn-hangzhou
	OutputLocation *string `json:"OutputLocation,omitempty" xml:"OutputLocation,omitempty"`
	// The job output configurations. For more information, see the "Output" section of the [Parameter details](https://help.aliyun.com/document_detail/29253.html) topic.
	//
	// 	- Specify the value in a JSON array of Output objects. You can specify up to 30 Output objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"OutputObject":"exampleOutput.mp4","TemplateId":"6181666213ab41b9bc21da8ff5ff****","WaterMarks":[{"InputFile":{"Bucket":"exampleBucket","Location":"oss-cn-hangzhou","Object":"image_01.png"},"WaterMarkTemplateId":"9b772ce2740d4d55876d8b542d47****"}],"UserData":"testid-001"}]
	Outputs      *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the MPS queue. For more information, see [Terms](https://help.aliyun.com/document_detail/31827.html).
	//
	// 	- To obtain the ID of an MPS queue, you can log on to the [MPS console](https://mps.console.aliyun.com/overview) and choose **Global Settings*	- > **MPS Queue and Callback*	- in the left-side navigation pane.
	//
	// 	- If you want to receive asynchronous message notifications, associate an MNS queue or topic with the MPS queue. For more information, see [Receive notifications](https://help.aliyun.com/document_detail/42618.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// dd3dae411e704030b921e52698e5****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SubmitJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitJobsRequest) GoString() string {
	return s.String()
}

func (s *SubmitJobsRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitJobsRequest) GetOutputBucket() *string {
	return s.OutputBucket
}

func (s *SubmitJobsRequest) GetOutputLocation() *string {
	return s.OutputLocation
}

func (s *SubmitJobsRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *SubmitJobsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitJobsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitJobsRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitJobsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitJobsRequest) SetInput(v string) *SubmitJobsRequest {
	s.Input = &v
	return s
}

func (s *SubmitJobsRequest) SetOutputBucket(v string) *SubmitJobsRequest {
	s.OutputBucket = &v
	return s
}

func (s *SubmitJobsRequest) SetOutputLocation(v string) *SubmitJobsRequest {
	s.OutputLocation = &v
	return s
}

func (s *SubmitJobsRequest) SetOutputs(v string) *SubmitJobsRequest {
	s.Outputs = &v
	return s
}

func (s *SubmitJobsRequest) SetOwnerAccount(v string) *SubmitJobsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitJobsRequest) SetOwnerId(v int64) *SubmitJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitJobsRequest) SetPipelineId(v string) *SubmitJobsRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitJobsRequest) SetResourceOwnerAccount(v string) *SubmitJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitJobsRequest) SetResourceOwnerId(v int64) *SubmitJobsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitJobsRequest) Validate() error {
	return dara.Validate(s)
}
