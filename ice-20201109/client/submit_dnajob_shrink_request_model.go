// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDNAJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *SubmitDNAJobShrinkRequest
	GetConfig() *string
	SetDBId(v string) *SubmitDNAJobShrinkRequest
	GetDBId() *string
	SetInputShrink(v string) *SubmitDNAJobShrinkRequest
	GetInputShrink() *string
	SetOwnerAccount(v string) *SubmitDNAJobShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SubmitDNAJobShrinkRequest
	GetOwnerId() *int64
	SetPipelineId(v string) *SubmitDNAJobShrinkRequest
	GetPipelineId() *string
	SetPrimaryKey(v string) *SubmitDNAJobShrinkRequest
	GetPrimaryKey() *string
	SetResourceOwnerAccount(v string) *SubmitDNAJobShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitDNAJobShrinkRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v string) *SubmitDNAJobShrinkRequest
	GetTemplateId() *string
	SetUserData(v string) *SubmitDNAJobShrinkRequest
	GetUserData() *string
}

type SubmitDNAJobShrinkRequest struct {
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
	InputShrink  *string `json:"Input,omitempty" xml:"Input,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s SubmitDNAJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDNAJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDNAJobShrinkRequest) GetConfig() *string {
	return s.Config
}

func (s *SubmitDNAJobShrinkRequest) GetDBId() *string {
	return s.DBId
}

func (s *SubmitDNAJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitDNAJobShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SubmitDNAJobShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitDNAJobShrinkRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitDNAJobShrinkRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *SubmitDNAJobShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitDNAJobShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitDNAJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitDNAJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitDNAJobShrinkRequest) SetConfig(v string) *SubmitDNAJobShrinkRequest {
	s.Config = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetDBId(v string) *SubmitDNAJobShrinkRequest {
	s.DBId = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetInputShrink(v string) *SubmitDNAJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetOwnerAccount(v string) *SubmitDNAJobShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetOwnerId(v int64) *SubmitDNAJobShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetPipelineId(v string) *SubmitDNAJobShrinkRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetPrimaryKey(v string) *SubmitDNAJobShrinkRequest {
	s.PrimaryKey = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetResourceOwnerAccount(v string) *SubmitDNAJobShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetResourceOwnerId(v int64) *SubmitDNAJobShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetTemplateId(v string) *SubmitDNAJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) SetUserData(v string) *SubmitDNAJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitDNAJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
