// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *GetApiKeyResponseBody
	GetApiKey() *string
	SetAuthServices(v []*GetApiKeyResponseBodyAuthServices) *GetApiKeyResponseBody
	GetAuthServices() []*GetApiKeyResponseBodyAuthServices
	SetCreateTime(v string) *GetApiKeyResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetApiKeyResponseBody
	GetDescription() *string
	SetKeyId(v string) *GetApiKeyResponseBody
	GetKeyId() *string
	SetKeyName(v string) *GetApiKeyResponseBody
	GetKeyName() *string
	SetKeyPrefix(v string) *GetApiKeyResponseBody
	GetKeyPrefix() *string
	SetRequestId(v string) *GetApiKeyResponseBody
	GetRequestId() *string
}

type GetApiKeyResponseBody struct {
	// The content of the API key.
	//
	// example:
	//
	// sk-xxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The service IDs.
	AuthServices []*GetApiKeyResponseBodyAuthServices `json:"AuthServices,omitempty" xml:"AuthServices,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2019-09-08T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test api key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API key.
	//
	// example:
	//
	// api-xxxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// test api key
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The prefix of the API key.
	//
	// example:
	//
	// sk-1235*****
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7565770E-7C45-462D-BA4A-8A5396****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetApiKeyResponseBody) GetAuthServices() []*GetApiKeyResponseBodyAuthServices {
	return s.AuthServices
}

func (s *GetApiKeyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApiKeyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetApiKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GetApiKeyResponseBody) GetKeyName() *string {
	return s.KeyName
}

func (s *GetApiKeyResponseBody) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *GetApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiKeyResponseBody) SetApiKey(v string) *GetApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *GetApiKeyResponseBody) SetAuthServices(v []*GetApiKeyResponseBodyAuthServices) *GetApiKeyResponseBody {
	s.AuthServices = v
	return s
}

func (s *GetApiKeyResponseBody) SetCreateTime(v string) *GetApiKeyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetApiKeyResponseBody) SetDescription(v string) *GetApiKeyResponseBody {
	s.Description = &v
	return s
}

func (s *GetApiKeyResponseBody) SetKeyId(v string) *GetApiKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GetApiKeyResponseBody) SetKeyName(v string) *GetApiKeyResponseBody {
	s.KeyName = &v
	return s
}

func (s *GetApiKeyResponseBody) SetKeyPrefix(v string) *GetApiKeyResponseBody {
	s.KeyPrefix = &v
	return s
}

func (s *GetApiKeyResponseBody) SetRequestId(v string) *GetApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiKeyResponseBody) Validate() error {
	if s.AuthServices != nil {
		for _, item := range s.AuthServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApiKeyResponseBodyAuthServices struct {
	// The service IDs.
	//
	// example:
	//
	// agdb-2ze8x9278c9iizl
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service type. Valid values:
	//
	// - **drama**
	//
	// - **memroy**
	//
	// example:
	//
	// memory
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetApiKeyResponseBodyAuthServices) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBodyAuthServices) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBodyAuthServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetApiKeyResponseBodyAuthServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetApiKeyResponseBodyAuthServices) SetServiceId(v string) *GetApiKeyResponseBodyAuthServices {
	s.ServiceId = &v
	return s
}

func (s *GetApiKeyResponseBodyAuthServices) SetServiceType(v string) *GetApiKeyResponseBodyAuthServices {
	s.ServiceType = &v
	return s
}

func (s *GetApiKeyResponseBodyAuthServices) Validate() error {
	return dara.Validate(s)
}
