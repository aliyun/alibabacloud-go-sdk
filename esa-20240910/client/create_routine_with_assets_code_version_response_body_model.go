// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineWithAssetsCodeVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *CreateRoutineWithAssetsCodeVersionResponseBody
	GetCodeVersion() *string
	SetOssPostConfig(v *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) *CreateRoutineWithAssetsCodeVersionResponseBody
	GetOssPostConfig() *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig
	SetRequestId(v string) *CreateRoutineWithAssetsCodeVersionResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateRoutineWithAssetsCodeVersionResponseBody
	GetStatus() *string
}

type CreateRoutineWithAssetsCodeVersionResponseBody struct {
	CodeVersion   *string                                                      `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	OssPostConfig *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig `json:"OssPostConfig,omitempty" xml:"OssPostConfig,omitempty" type:"Struct"`
	RequestId     *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRoutineWithAssetsCodeVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineWithAssetsCodeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) GetOssPostConfig() *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig {
	return s.OssPostConfig
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) SetCodeVersion(v string) *CreateRoutineWithAssetsCodeVersionResponseBody {
	s.CodeVersion = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) SetOssPostConfig(v *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) *CreateRoutineWithAssetsCodeVersionResponseBody {
	s.OssPostConfig = v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) SetRequestId(v string) *CreateRoutineWithAssetsCodeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) SetStatus(v string) *CreateRoutineWithAssetsCodeVersionResponseBody {
	s.Status = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig struct {
	Key               *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OSSAccessKeyId    *string `json:"OSSAccessKeyId,omitempty" xml:"OSSAccessKeyId,omitempty"`
	Policy            *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Signature         *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Url               *string `json:"Url,omitempty" xml:"Url,omitempty"`
	XOssSecurityToken *string `json:"XOssSecurityToken,omitempty" xml:"XOssSecurityToken,omitempty"`
}

func (s CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) GoString() string {
	return s.String()
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) GetKey() *string {
	return s.Key
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) GetOSSAccessKeyId() *string {
	return s.OSSAccessKeyId
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) GetPolicy() *string {
	return s.Policy
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) GetSignature() *string {
	return s.Signature
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) GetUrl() *string {
	return s.Url
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) GetXOssSecurityToken() *string {
	return s.XOssSecurityToken
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) SetKey(v string) *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig {
	s.Key = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) SetOSSAccessKeyId(v string) *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig {
	s.OSSAccessKeyId = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) SetPolicy(v string) *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig {
	s.Policy = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) SetSignature(v string) *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig {
	s.Signature = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) SetUrl(v string) *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig {
	s.Url = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) SetXOssSecurityToken(v string) *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig {
	s.XOssSecurityToken = &v
	return s
}

func (s *CreateRoutineWithAssetsCodeVersionResponseBodyOssPostConfig) Validate() error {
	return dara.Validate(s)
}
