// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDNAJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *SubmitDNAJobRequest
	GetConfig() *string
	SetDBId(v string) *SubmitDNAJobRequest
	GetDBId() *string
	SetInput(v *SubmitDNAJobRequestInput) *SubmitDNAJobRequest
	GetInput() *SubmitDNAJobRequestInput
	SetOwnerAccount(v string) *SubmitDNAJobRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitDNAJobRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitDNAJobRequest
	GetPipelineId() *string
	SetPrimaryKey(v string) *SubmitDNAJobRequest
	GetPrimaryKey() *string
	SetResourceOwnerAccount(v string) *SubmitDNAJobRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitDNAJobRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v string) *SubmitDNAJobRequest
	GetTemplateId() *string
	SetUserData(v string) *SubmitDNAJobRequest
	GetUserData() *string
}

type SubmitDNAJobRequest struct {
	// The configurations of the media fingerprint analysis job. The value is a JSON object. If you specify this parameter, the template parameters are overwritten.
	//
	// example:
	//
	// {"SaveType": "save","MediaType"":"video"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the media fingerprint library. If you do not specify this parameter, the default media fingerprint library is used. For more information about how to create a media fingerprint library, see [CreateDNADB](https://help.aliyun.com/document_detail/479275.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	DBId *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	// The input file for media fingerprint analysis.
	//
	// This parameter is required.
	Input        *SubmitDNAJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	OwnerAccount *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the ApsaraVideo Media Processing (MPS) queue to which the media fingerprint analysis job is submitted.
	//
	// example:
	//
	// 5246b8d12a62433ab77845074039****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The primary key of the video. You must make sure that each primary key is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3ca84a39a9024f19853b21be9cf9****
	PrimaryKey           *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template ID.
	//
	// example:
	//
	// S00000101-100060
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The user-defined data. The data can be up to 128 bytes in length.
	//
	// example:
	//
	// userData
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitDNAJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDNAJobRequest) GetConfig() *string {
	return s.Config
}

func (s *SubmitDNAJobRequest) GetDBId() *string {
	return s.DBId
}

func (s *SubmitDNAJobRequest) GetInput() *SubmitDNAJobRequestInput {
	return s.Input
}

func (s *SubmitDNAJobRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitDNAJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitDNAJobRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitDNAJobRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *SubmitDNAJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitDNAJobRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitDNAJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitDNAJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitDNAJobRequest) SetConfig(v string) *SubmitDNAJobRequest {
	s.Config = &v
	return s
}

func (s *SubmitDNAJobRequest) SetDBId(v string) *SubmitDNAJobRequest {
	s.DBId = &v
	return s
}

func (s *SubmitDNAJobRequest) SetInput(v *SubmitDNAJobRequestInput) *SubmitDNAJobRequest {
	s.Input = v
	return s
}

func (s *SubmitDNAJobRequest) SetOwnerAccount(v string) *SubmitDNAJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitDNAJobRequest) SetOwnerId(v int64) *SubmitDNAJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitDNAJobRequest) SetPipelineId(v string) *SubmitDNAJobRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitDNAJobRequest) SetPrimaryKey(v string) *SubmitDNAJobRequest {
	s.PrimaryKey = &v
	return s
}

func (s *SubmitDNAJobRequest) SetResourceOwnerAccount(v string) *SubmitDNAJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitDNAJobRequest) SetResourceOwnerId(v int64) *SubmitDNAJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitDNAJobRequest) SetTemplateId(v string) *SubmitDNAJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitDNAJobRequest) SetUserData(v string) *SubmitDNAJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitDNAJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDNAJobRequestInput struct {
	// The input file. The file can be an OSS object or a media asset. You can specify the path of an OSS object in one of the following formats:
	//
	// 1\\. oss://bucket/object
	//
	// 2\\. http(s)://bucket.oss-[regionId].aliyuncs.com/object
	//
	// In the preceding paths, bucket indicates an OSS bucket that resides in the same region as the current project, and object indicates the path of the object in the bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1b1b9cd148034739af413150fded****
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1.  OSS: Object Storage Service (OSS) object.
	//
	// 2.  Media: media asset.
	//
	// This parameter is required.
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitDNAJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitDNAJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitDNAJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitDNAJobRequestInput) SetMedia(v string) *SubmitDNAJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitDNAJobRequestInput) SetType(v string) *SubmitDNAJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitDNAJobRequestInput) Validate() error {
	return dara.Validate(s)
}
