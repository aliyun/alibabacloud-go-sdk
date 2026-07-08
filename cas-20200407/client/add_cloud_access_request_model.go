// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudName(v string) *AddCloudAccessRequest
	GetCloudName() *string
	SetSecretId(v string) *AddCloudAccessRequest
	GetSecretId() *string
	SetSecretKey(v string) *AddCloudAccessRequest
	GetSecretKey() *string
}

type AddCloudAccessRequest struct {
	// The cloud service provider. This API supports multiple providers as detailed in the SecretKey parameter description. For example, to add credentials for Tencent Cloud, set this parameter to **Tencent**.
	//
	// example:
	//
	// Tencent
	CloudName *string `json:"CloudName,omitempty" xml:"CloudName,omitempty"`
	// The Secret ID for accessing the cloud resource set.
	//
	// example:
	//
	// xcxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The secret corresponding to the AccessKey. The value is determined by the `AkType` parameter as follows:
	//
	// 1\\. If `AkType` is set to `primary`:
	//
	// - **Tencent**: The SecretAccessKey of the primary account.
	//
	// - **HUAWEI CLOUD**: The SecretAccessKey of the primary account.
	//
	// - **Azure**: The ClientSecret.
	//
	// - **AWS**: The SecretAccessKey of the primary account.
	//
	// 2\\. If `AkType` is set to `sub`:
	//
	// - **Tencent**: The SecretAccessKey of the sub-account.
	//
	// - **HUAWEI CLOUD**: The SecretAccessKey of the sub-account.
	//
	// - **Azure**: The ClientSecret.
	//
	// - **AWS**: The SecretAccessKey of the sub-account.
	//
	// example:
	//
	// xxx
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
}

func (s AddCloudAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCloudAccessRequest) GoString() string {
	return s.String()
}

func (s *AddCloudAccessRequest) GetCloudName() *string {
	return s.CloudName
}

func (s *AddCloudAccessRequest) GetSecretId() *string {
	return s.SecretId
}

func (s *AddCloudAccessRequest) GetSecretKey() *string {
	return s.SecretKey
}

func (s *AddCloudAccessRequest) SetCloudName(v string) *AddCloudAccessRequest {
	s.CloudName = &v
	return s
}

func (s *AddCloudAccessRequest) SetSecretId(v string) *AddCloudAccessRequest {
	s.SecretId = &v
	return s
}

func (s *AddCloudAccessRequest) SetSecretKey(v string) *AddCloudAccessRequest {
	s.SecretKey = &v
	return s
}

func (s *AddCloudAccessRequest) Validate() error {
	return dara.Validate(s)
}
