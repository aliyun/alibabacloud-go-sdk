// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportFpShotJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFpDBId(v string) *ImportFpShotJobRequest
	GetFpDBId() *string
	SetFpImportConfig(v string) *ImportFpShotJobRequest
	GetFpImportConfig() *string
	SetInput(v string) *ImportFpShotJobRequest
	GetInput() *string
	SetOwnerAccount(v string) *ImportFpShotJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ImportFpShotJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *ImportFpShotJobRequest
	GetPipelineId() *string
	SetResourceOwnerAccount(v string) *ImportFpShotJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportFpShotJobRequest
	GetResourceOwnerId() *int64
	SetUserData(v string) *ImportFpShotJobRequest
	GetUserData() *string
}

type ImportFpShotJobRequest struct {
	// The ID of the text fingerprint library to which the text file is imported. You can specify only one job of importing text files to a text fingerprint library at a time. You can obtain the library ID from the response parameters of the [CreateFpShotDB](https://help.aliyun.com/document_detail/170149.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	FpDBId *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	// The job configurations. The value must be a JSON object. Example: `{"SaveType":"onlysave"}`. The `SaveType` field indicates the storage type. Valid values of the SaveType field:
	//
	// 	- **save**: The fingerprints of the text file are saved to the text fingerprint library only if the text file is not duplicated with content in the text fingerprint library.
	//
	// 	- **onlysave**: The fingerprints of the text file are saved to the text fingerprint library.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"SaveType":"onlysave"}
	FpImportConfig *string `json:"FpImportConfig,omitempty" xml:"FpImportConfig,omitempty"`
	// The Object Storage Service (OSS) URL of the text file to be imported to the text fingerprint library. The value must be a JSON object. Example: {"Bucket":"example-bucket","Location":"oss-cn-shanghai","Object":"example.flv"}.
	//
	// > The OSS bucket must reside in the same region as your MPS service.
	//
	// This parameter is required.
	//
	// example:
	//
	// {“Bucket”:”example-bucket”,“Location”:”oss-cn-shanghai”,“Object”:”example.txt”}
	Input        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue. To view the ID of the MPS queue, perform the following steps: Log on to the **MPS console**. In the left-side navigation pane, choose **Global Settings*	- > **Pipelines**. The MPS queue is associated with a specified Message Service (MNS) topic. You can submit jobs for different services to different MPS queues. If you do not specify this parameter, the job is submitted to the default MPS queue and no MNS topic is associated with the MPS queue.
	//
	// example:
	//
	// ae687c02fe944327ba9631e50da2****
	PipelineId           *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The user-defined data. The value can contain letters, digits, and special characters. The value can be up to 128 bytes in length.
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ImportFpShotJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportFpShotJobRequest) GoString() string {
	return s.String()
}

func (s *ImportFpShotJobRequest) GetFpDBId() *string {
	return s.FpDBId
}

func (s *ImportFpShotJobRequest) GetFpImportConfig() *string {
	return s.FpImportConfig
}

func (s *ImportFpShotJobRequest) GetInput() *string {
	return s.Input
}

func (s *ImportFpShotJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ImportFpShotJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportFpShotJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ImportFpShotJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportFpShotJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportFpShotJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *ImportFpShotJobRequest) SetFpDBId(v string) *ImportFpShotJobRequest {
	s.FpDBId = &v
	return s
}

func (s *ImportFpShotJobRequest) SetFpImportConfig(v string) *ImportFpShotJobRequest {
	s.FpImportConfig = &v
	return s
}

func (s *ImportFpShotJobRequest) SetInput(v string) *ImportFpShotJobRequest {
	s.Input = &v
	return s
}

func (s *ImportFpShotJobRequest) SetOwnerAccount(v string) *ImportFpShotJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ImportFpShotJobRequest) SetOwnerId(v int64) *ImportFpShotJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportFpShotJobRequest) SetPipelineId(v string) *ImportFpShotJobRequest {
	s.PipelineId = &v
	return s
}

func (s *ImportFpShotJobRequest) SetResourceOwnerAccount(v string) *ImportFpShotJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportFpShotJobRequest) SetResourceOwnerId(v int64) *ImportFpShotJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportFpShotJobRequest) SetUserData(v string) *ImportFpShotJobRequest {
	s.UserData = &v
	return s
}

func (s *ImportFpShotJobRequest) Validate() error {
	return dara.Validate(s)
}
