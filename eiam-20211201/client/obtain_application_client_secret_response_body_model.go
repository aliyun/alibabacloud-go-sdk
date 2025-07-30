// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainApplicationClientSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationClientSecret(v *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) *ObtainApplicationClientSecretResponseBody
	GetApplicationClientSecret() *ObtainApplicationClientSecretResponseBodyApplicationClientSecret
	SetRequestId(v string) *ObtainApplicationClientSecretResponseBody
	GetRequestId() *string
}

type ObtainApplicationClientSecretResponseBody struct {
	// The information about the client key.
	ApplicationClientSecret *ObtainApplicationClientSecretResponseBodyApplicationClientSecret `json:"ApplicationClientSecret,omitempty" xml:"ApplicationClientSecret,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ObtainApplicationClientSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretResponseBody) GetApplicationClientSecret() *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	return s.ApplicationClientSecret
}

func (s *ObtainApplicationClientSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ObtainApplicationClientSecretResponseBody) SetApplicationClientSecret(v *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) *ObtainApplicationClientSecretResponseBody {
	s.ApplicationClientSecret = v
	return s
}

func (s *ObtainApplicationClientSecretResponseBody) SetRequestId(v string) *ObtainApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBody) Validate() error {
	return dara.Validate(s)
}

type ObtainApplicationClientSecretResponseBodyApplicationClientSecret struct {
	// The ID of the application whose client key you want to query.
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
	// The client key secret of the application.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
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

func (s ObtainApplicationClientSecretResponseBodyApplicationClientSecret) String() string {
	return dara.Prettify(s)
}

func (s ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GetClientId() *string {
	return s.ClientId
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GetLastUsedTime() *int64 {
	return s.LastUsedTime
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GetSecretId() *string {
	return s.SecretId
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GetStatus() *string {
	return s.Status
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetApplicationId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ApplicationId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetClientId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetClientSecret(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientSecret = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetInstanceId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.InstanceId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetLastUsedTime(v int64) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.LastUsedTime = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetSecretId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.SecretId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetStatus(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.Status = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) Validate() error {
	return dara.Validate(s)
}
