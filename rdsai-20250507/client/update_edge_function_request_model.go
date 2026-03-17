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
	SetCustomConfig(v map[string]interface{}) *UpdateEdgeFunctionRequest
	GetCustomConfig() map[string]interface{}
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
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Code        *UpdateEdgeFunctionRequestCode `json:"Code,omitempty" xml:"Code,omitempty" type:"Struct"`
	// example:
	//
	// {}
	CustomConfig map[string]interface{} `json:"CustomConfig,omitempty" xml:"CustomConfig,omitempty"`
	// fc-xxxx。
	//
	// example:
	//
	// ef-****
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

func (s *UpdateEdgeFunctionRequest) GetCustomConfig() map[string]interface{} {
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

func (s *UpdateEdgeFunctionRequest) SetCustomConfig(v map[string]interface{}) *UpdateEdgeFunctionRequest {
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
	// example:
	//
	// code
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// example2.zip
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
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

func (s *UpdateEdgeFunctionRequestCode) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *UpdateEdgeFunctionRequestCode) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *UpdateEdgeFunctionRequestCode) GetOssType() *string {
	return s.OssType
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
