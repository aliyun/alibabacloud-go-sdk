// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEdgeFunctionRequest
	GetClientToken() *string
	SetCode(v *CreateEdgeFunctionRequestCode) *CreateEdgeFunctionRequest
	GetCode() *CreateEdgeFunctionRequestCode
	SetCustomConfig(v map[string]*int32) *CreateEdgeFunctionRequest
	GetCustomConfig() map[string]*int32
	SetEdgeFunctionName(v string) *CreateEdgeFunctionRequest
	GetEdgeFunctionName() *string
	SetEnvs(v map[string]*string) *CreateEdgeFunctionRequest
	GetEnvs() map[string]*string
	SetInstanceName(v string) *CreateEdgeFunctionRequest
	GetInstanceName() *string
	SetRegionId(v string) *CreateEdgeFunctionRequest
	GetRegionId() *string
}

type CreateEdgeFunctionRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken  *string                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Code         *CreateEdgeFunctionRequestCode `json:"Code,omitempty" xml:"Code,omitempty" type:"Struct"`
	CustomConfig map[string]*int32              `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// example:
	//
	// ef-*****
	EdgeFunctionName *string            `json:"EdgeFunctionName,omitempty" xml:"EdgeFunctionName,omitempty"`
	Envs             map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEdgeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeFunctionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEdgeFunctionRequest) GetCode() *CreateEdgeFunctionRequestCode {
	return s.Code
}

func (s *CreateEdgeFunctionRequest) GetCustomConfig() map[string]*int32 {
	return s.CustomConfig
}

func (s *CreateEdgeFunctionRequest) GetEdgeFunctionName() *string {
	return s.EdgeFunctionName
}

func (s *CreateEdgeFunctionRequest) GetEnvs() map[string]*string {
	return s.Envs
}

func (s *CreateEdgeFunctionRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateEdgeFunctionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEdgeFunctionRequest) SetClientToken(v string) *CreateEdgeFunctionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEdgeFunctionRequest) SetCode(v *CreateEdgeFunctionRequestCode) *CreateEdgeFunctionRequest {
	s.Code = v
	return s
}

func (s *CreateEdgeFunctionRequest) SetCustomConfig(v map[string]*int32) *CreateEdgeFunctionRequest {
	s.CustomConfig = v
	return s
}

func (s *CreateEdgeFunctionRequest) SetEdgeFunctionName(v string) *CreateEdgeFunctionRequest {
	s.EdgeFunctionName = &v
	return s
}

func (s *CreateEdgeFunctionRequest) SetEnvs(v map[string]*string) *CreateEdgeFunctionRequest {
	s.Envs = v
	return s
}

func (s *CreateEdgeFunctionRequest) SetInstanceName(v string) *CreateEdgeFunctionRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateEdgeFunctionRequest) SetRegionId(v string) *CreateEdgeFunctionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEdgeFunctionRequest) Validate() error {
	if s.Code != nil {
		if err := s.Code.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEdgeFunctionRequestCode struct {
	// example:
	//
	// code
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// example.zip
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// supabase
	OssType *string `json:"OssType,omitempty" xml:"OssType,omitempty"`
}

func (s CreateEdgeFunctionRequestCode) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeFunctionRequestCode) GoString() string {
	return s.String()
}

func (s *CreateEdgeFunctionRequestCode) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CreateEdgeFunctionRequestCode) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *CreateEdgeFunctionRequestCode) GetOssType() *string {
	return s.OssType
}

func (s *CreateEdgeFunctionRequestCode) SetOssBucketName(v string) *CreateEdgeFunctionRequestCode {
	s.OssBucketName = &v
	return s
}

func (s *CreateEdgeFunctionRequestCode) SetOssObjectName(v string) *CreateEdgeFunctionRequestCode {
	s.OssObjectName = &v
	return s
}

func (s *CreateEdgeFunctionRequestCode) SetOssType(v string) *CreateEdgeFunctionRequestCode {
	s.OssType = &v
	return s
}

func (s *CreateEdgeFunctionRequestCode) Validate() error {
	return dara.Validate(s)
}
