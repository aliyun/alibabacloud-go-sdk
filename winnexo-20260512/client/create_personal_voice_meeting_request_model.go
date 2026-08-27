// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalVoiceMeetingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalVoiceMeetingRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalVoiceMeetingRequest
	GetDirectoryId() *string
	SetFileUrl(v string) *CreatePersonalVoiceMeetingRequest
	GetFileUrl() *string
	SetName(v string) *CreatePersonalVoiceMeetingRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreatePersonalVoiceMeetingRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *CreatePersonalVoiceMeetingRequest
	GetTenantId() *string
}

type CreatePersonalVoiceMeetingRequest struct {
	// The pipeline description.
	//
	// example:
	//
	// update-JNQ9csEc6ArPPSXANH7O
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The Yida attachment URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://emas-devops-cdn.aliyuncs.com/job/PJ859733644657824768/apk/release/com.czmp.vitanexusoff_release_v331.178.14-signed.apk
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// p-toolset-2a1461ff-59c1-4baa-9e19-966ec7c00004
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 474379246158592
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalVoiceMeetingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalVoiceMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalVoiceMeetingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalVoiceMeetingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalVoiceMeetingRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreatePersonalVoiceMeetingRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalVoiceMeetingRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalVoiceMeetingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalVoiceMeetingRequest) SetDescription(v string) *CreatePersonalVoiceMeetingRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetDirectoryId(v string) *CreatePersonalVoiceMeetingRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetFileUrl(v string) *CreatePersonalVoiceMeetingRequest {
	s.FileUrl = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetName(v string) *CreatePersonalVoiceMeetingRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetOperatingObjectName(v string) *CreatePersonalVoiceMeetingRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) SetTenantId(v string) *CreatePersonalVoiceMeetingRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingRequest) Validate() error {
	return dara.Validate(s)
}
