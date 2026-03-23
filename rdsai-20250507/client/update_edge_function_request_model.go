// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateEdgeFunctionRequest
	GetClientToken() *string
	SetCode(v *UpdateEdgeFunctionRequestCode) *UpdateEdgeFunctionRequest
	GetCode() *UpdateEdgeFunctionRequestCode
	SetCustomConfig(v map[string]*string) *UpdateEdgeFunctionRequest
	GetCustomConfig() map[string]*string
	SetEdgeFunctionName(v string) *UpdateEdgeFunctionRequest
	GetEdgeFunctionName() *string
	SetEnvs(v map[string]*string) *UpdateEdgeFunctionRequest
	GetEnvs() map[string]*string
	SetInstanceName(v string) *UpdateEdgeFunctionRequest
	GetInstanceName() *string
	SetRegionId(v string) *UpdateEdgeFunctionRequest
	GetRegionId() *string
}

type UpdateEdgeFunctionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The error code that is returned if the request failed. For more information, see the "Error codes" section of the topic.
	Code *UpdateEdgeFunctionRequestCode `json:"Code,omitempty" xml:"Code,omitempty" type:"Struct"`
	// The configuration parameters of the edge function.
	//
	// example:
	//
	// {}
	CustomConfig map[string]*string `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// fc-xxxx
	//
	// example:
	//
	// ef-****
	EdgeFunctionName *string `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	// The environment variables of the edge function.
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The ID of the RDS Supabase instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEdgeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateEdgeFunctionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateEdgeFunctionRequest) GetCode() *UpdateEdgeFunctionRequestCode {
	return s.Code
}

func (s *UpdateEdgeFunctionRequest) GetCustomConfig() map[string]*string {
	return s.CustomConfig
}

func (s *UpdateEdgeFunctionRequest) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *UpdateEdgeFunctionRequest) GetEnvs() map[string]*string {
	return s.Envs
}

func (s *UpdateEdgeFunctionRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateEdgeFunctionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEdgeFunctionRequest) SetClientToken(v string) *UpdateEdgeFunctionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateEdgeFunctionRequest) SetCode(v *UpdateEdgeFunctionRequestCode) *UpdateEdgeFunctionRequest {
	s.Code = v
	return s
}

func (s *UpdateEdgeFunctionRequest) SetCustomConfig(v map[string]*string) *UpdateEdgeFunctionRequest {
	s.CustomConfig = v
	return s
}

func (s *UpdateEdgeFunctionRequest) SetEdgeFunctionName(v string) *UpdateEdgeFunctionRequest {
	s.EdgeFunctionName = &v
	return s
}

func (s *UpdateEdgeFunctionRequest) SetEnvs(v map[string]*string) *UpdateEdgeFunctionRequest {
	s.Envs = v
	return s
}

func (s *UpdateEdgeFunctionRequest) SetInstanceName(v string) *UpdateEdgeFunctionRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateEdgeFunctionRequest) SetRegionId(v string) *UpdateEdgeFunctionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEdgeFunctionRequest) Validate() error {
	if s.Code != nil {
		if err := s.Code.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateEdgeFunctionRequestCode struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// code
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The path of the code file.
	//
	// example:
	//
	// example2.zip
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// The storage class of the OSS bucket.
	//
	// example:
	//
	// supabase
	OssType *string `json:"OssType,omitempty" xml:"OssType,omitempty"`
}

func (s UpdateEdgeFunctionRequestCode) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeFunctionRequestCode) GoString() string {
	return s.String()
}

func (s *UpdateEdgeFunctionRequestCode) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *UpdateEdgeFunctionRequestCode) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *UpdateEdgeFunctionRequestCode) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *UpdateEdgeFunctionRequestCode) GetOssType() *string {
	return s.OssType
}

func (s *UpdateEdgeFunctionRequestCode) SetDownloadUrl(v string) *UpdateEdgeFunctionRequestCode {
	s.DownloadUrl = &v
	return s
}

func (s *UpdateEdgeFunctionRequestCode) SetOssBucketName(v string) *UpdateEdgeFunctionRequestCode {
	s.OssBucketName = &v
	return s
}

func (s *UpdateEdgeFunctionRequestCode) SetOssObjectName(v string) *UpdateEdgeFunctionRequestCode {
	s.OssObjectName = &v
	return s
}

func (s *UpdateEdgeFunctionRequestCode) SetOssType(v string) *UpdateEdgeFunctionRequestCode {
	s.OssType = &v
	return s
}

func (s *UpdateEdgeFunctionRequestCode) Validate() error {
	return dara.Validate(s)
}
