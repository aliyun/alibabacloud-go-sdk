// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *UploadInfo
	GetAccessId() *string
	SetHost(v string) *UploadInfo
	GetHost() *string
	SetKey(v string) *UploadInfo
	GetKey() *string
	SetPolicy(v string) *UploadInfo
	GetPolicy() *string
	SetPolicySignature(v string) *UploadInfo
	GetPolicySignature() *string
	SetUrl(v string) *UploadInfo
	GetUrl() *string
}

type UploadInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxxx
	AccessId *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yic-pre.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234/temp-novels/xxxx-xxx-xx.txt
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxxx
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxxxxx
	PolicySignature *string `json:"policySignature,omitempty" xml:"policySignature,omitempty"`
	// example:
	//
	// xxxxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UploadInfo) String() string {
	return dara.Prettify(s)
}

func (s UploadInfo) GoString() string {
	return s.String()
}

func (s *UploadInfo) GetAccessId() *string {
	return s.AccessId
}

func (s *UploadInfo) GetHost() *string {
	return s.Host
}

func (s *UploadInfo) GetKey() *string {
	return s.Key
}

func (s *UploadInfo) GetPolicy() *string {
	return s.Policy
}

func (s *UploadInfo) GetPolicySignature() *string {
	return s.PolicySignature
}

func (s *UploadInfo) GetUrl() *string {
	return s.Url
}

func (s *UploadInfo) SetAccessId(v string) *UploadInfo {
	s.AccessId = &v
	return s
}

func (s *UploadInfo) SetHost(v string) *UploadInfo {
	s.Host = &v
	return s
}

func (s *UploadInfo) SetKey(v string) *UploadInfo {
	s.Key = &v
	return s
}

func (s *UploadInfo) SetPolicy(v string) *UploadInfo {
	s.Policy = &v
	return s
}

func (s *UploadInfo) SetPolicySignature(v string) *UploadInfo {
	s.PolicySignature = &v
	return s
}

func (s *UploadInfo) SetUrl(v string) *UploadInfo {
	s.Url = &v
	return s
}

func (s *UploadInfo) Validate() error {
	return dara.Validate(s)
}
