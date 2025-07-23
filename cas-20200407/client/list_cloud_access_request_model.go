// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudName(v string) *ListCloudAccessRequest
	GetCloudName() *string
	SetCurrentPage(v int32) *ListCloudAccessRequest
	GetCurrentPage() *int32
	SetSecretId(v string) *ListCloudAccessRequest
	GetSecretId() *string
	SetShowSize(v int32) *ListCloudAccessRequest
	GetShowSize() *int32
}

type ListCloudAccessRequest struct {
	// The cloud service provider. Set the value to **Tencent**, which indicates Tencent Cloud.
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The AccessKey ID that is used to access cloud resources.
	//
	// example:
	//
	// 276
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The number of entries per page. Valid values: **10**, **20**, and **50**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCloudAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccessRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAccessRequest) GetCloudName() *string {
	return s.CloudName
}

func (s *ListCloudAccessRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudAccessRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *ListCloudAccessRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCloudAccessRequest) SetCloudName(v string) *ListCloudAccessRequest {
	s.CloudName = &v
	return s
}

func (s *ListCloudAccessRequest) SetCurrentPage(v int32) *ListCloudAccessRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAccessRequest) SetSecretId(v string) *ListCloudAccessRequest {
	s.SecretId = &v
	return s
}

func (s *ListCloudAccessRequest) SetShowSize(v int32) *ListCloudAccessRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCloudAccessRequest) Validate() error {
	return dara.Validate(s)
}
