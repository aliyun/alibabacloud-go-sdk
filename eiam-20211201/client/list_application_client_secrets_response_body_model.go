// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationClientSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationClientSecrets(v []*ListApplicationClientSecretsResponseBodyApplicationClientSecrets) *ListApplicationClientSecretsResponseBody
	GetApplicationClientSecrets() []*ListApplicationClientSecretsResponseBodyApplicationClientSecrets
	SetRequestId(v string) *ListApplicationClientSecretsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationClientSecretsResponseBody
	GetTotalCount() *int64
}

type ListApplicationClientSecretsResponseBody struct {
	// The information about the client keys.
	ApplicationClientSecrets []*ListApplicationClientSecretsResponseBodyApplicationClientSecrets `json:"ApplicationClientSecrets,omitempty" xml:"ApplicationClientSecrets,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationClientSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationClientSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsResponseBody) GetApplicationClientSecrets() []*ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	return s.ApplicationClientSecrets
}

func (s *ListApplicationClientSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationClientSecretsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationClientSecretsResponseBody) SetApplicationClientSecrets(v []*ListApplicationClientSecretsResponseBodyApplicationClientSecrets) *ListApplicationClientSecretsResponseBody {
	s.ApplicationClientSecrets = v
	return s
}

func (s *ListApplicationClientSecretsResponseBody) SetRequestId(v string) *ListApplicationClientSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBody) SetTotalCount(v int64) *ListApplicationClientSecretsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationClientSecretsResponseBodyApplicationClientSecrets struct {
	// The ID of the application that you want to query.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client key secret of the application. The value is not masked.
	//
	// example:
	//
	// eyJh*****************************************************************************************************OQ
	ClientSecret   *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	ExpirationTime *int64  `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_wdziy4vnjt33ehhf7z2o2nxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the client key was last used. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The client key ID of the application.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The status of the client key. Valid values:
	//
	// 	- Enabled: The client key is enabled.
	//
	// 	- Disabled: The client key is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationClientSecretsResponseBodyApplicationClientSecrets) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetClientId() *string {
	return s.ClientId
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetSecretId() *string {
	return s.SecretId
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GetStatus() *string {
	return s.Status
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetApplicationId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetClientId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.ClientId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetClientSecret(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.ClientSecret = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetExpirationTime(v int64) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.ExpirationTime = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetInstanceId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetLastUsedTime(v int64) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.LastUsedTime = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetSecretId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.SecretId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetStatus(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.Status = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) Validate() error {
	return dara.Validate(s)
}
