// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *GetDRMLicenseRequest
	GetKeyId() *string
	SetNotifyEndpoint(v string) *GetDRMLicenseRequest
	GetNotifyEndpoint() *string
	SetNotifyTopicName(v string) *GetDRMLicenseRequest
	GetNotifyTopicName() *string
	SetProjectName(v string) *GetDRMLicenseRequest
	GetProjectName() *string
	SetProtectionSystem(v string) *GetDRMLicenseRequest
	GetProtectionSystem() *string
}

type GetDRMLicenseRequest struct {
	// example:
	//
	// AESzB8SQgpACioSEJ3yqiFwruAOUgIvlCx*****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// example:
	//
	// http://1111111111.mns.cn-hangzhou.aliyuncs.com
	NotifyEndpoint *string `json:"NotifyEndpoint,omitempty" xml:"NotifyEndpoint,omitempty"`
	// example:
	//
	// topic1
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// widevine
	ProtectionSystem *string `json:"ProtectionSystem,omitempty" xml:"ProtectionSystem,omitempty"`
}

func (s GetDRMLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDRMLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GetDRMLicenseRequest) GetNotifyEndpoint() *string {
	return s.NotifyEndpoint
}

func (s *GetDRMLicenseRequest) GetNotifyTopicName() *string {
	return s.NotifyTopicName
}

func (s *GetDRMLicenseRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDRMLicenseRequest) GetProtectionSystem() *string {
	return s.ProtectionSystem
}

func (s *GetDRMLicenseRequest) SetKeyId(v string) *GetDRMLicenseRequest {
	s.KeyId = &v
	return s
}

func (s *GetDRMLicenseRequest) SetNotifyEndpoint(v string) *GetDRMLicenseRequest {
	s.NotifyEndpoint = &v
	return s
}

func (s *GetDRMLicenseRequest) SetNotifyTopicName(v string) *GetDRMLicenseRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GetDRMLicenseRequest) SetProjectName(v string) *GetDRMLicenseRequest {
	s.ProjectName = &v
	return s
}

func (s *GetDRMLicenseRequest) SetProtectionSystem(v string) *GetDRMLicenseRequest {
	s.ProtectionSystem = &v
	return s
}

func (s *GetDRMLicenseRequest) Validate() error {
	return dara.Validate(s)
}
