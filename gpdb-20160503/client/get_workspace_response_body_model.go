// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApikeys(v []*GetWorkspaceResponseBodyApikeys) *GetWorkspaceResponseBody
	GetApikeys() []*GetWorkspaceResponseBodyApikeys
	SetCreateTime(v string) *GetWorkspaceResponseBody
	GetCreateTime() *string
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetServices(v []*GetWorkspaceResponseBodyServices) *GetWorkspaceResponseBody
	GetServices() []*GetWorkspaceResponseBodyServices
	SetWorkspaceId(v string) *GetWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *GetWorkspaceResponseBody
	GetWorkspaceName() *string
}

type GetWorkspaceResponseBody struct {
	// The list of workspace API keys.
	Apikeys []*GetWorkspaceResponseBodyApikeys `json:"Apikeys,omitempty" xml:"Apikeys,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// 2026-06-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of service details.
	Services []*GetWorkspaceResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// my-first-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetApikeys() []*GetWorkspaceResponseBodyApikeys {
	return s.Apikeys
}

func (s *GetWorkspaceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetServices() []*GetWorkspaceResponseBodyServices {
	return s.Services
}

func (s *GetWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetWorkspaceResponseBody) SetApikeys(v []*GetWorkspaceResponseBodyApikeys) *GetWorkspaceResponseBody {
	s.Apikeys = v
	return s
}

func (s *GetWorkspaceResponseBody) SetCreateTime(v string) *GetWorkspaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetServices(v []*GetWorkspaceResponseBodyServices) *GetWorkspaceResponseBody {
	s.Services = v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceId(v string) *GetWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceName(v string) *GetWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	if s.Apikeys != nil {
		for _, item := range s.Apikeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkspaceResponseBodyApikeys struct {
	// The service ID.
	AuthServices []*GetWorkspaceResponseBodyApikeysAuthServices `json:"AuthServices,omitempty" xml:"AuthServices,omitempty" type:"Repeated"`
	// The creation time.
	//
	// example:
	//
	// Sat Mar 14 14:44:27 GMT+08:00 2026
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// my test key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API key.
	//
	// example:
	//
	// api-xxxxxx
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// my test key
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The prefix of the API key.
	//
	// example:
	//
	// sk-1235*****
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
}

func (s GetWorkspaceResponseBodyApikeys) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyApikeys) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyApikeys) GetAuthServices() []*GetWorkspaceResponseBodyApikeysAuthServices {
	return s.AuthServices
}

func (s *GetWorkspaceResponseBodyApikeys) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBodyApikeys) GetDescription() *string {
	return s.Description
}

func (s *GetWorkspaceResponseBodyApikeys) GetKeyId() *string {
	return s.KeyId
}

func (s *GetWorkspaceResponseBodyApikeys) GetKeyName() *string {
	return s.KeyName
}

func (s *GetWorkspaceResponseBodyApikeys) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *GetWorkspaceResponseBodyApikeys) SetAuthServices(v []*GetWorkspaceResponseBodyApikeysAuthServices) *GetWorkspaceResponseBodyApikeys {
	s.AuthServices = v
	return s
}

func (s *GetWorkspaceResponseBodyApikeys) SetCreateTime(v string) *GetWorkspaceResponseBodyApikeys {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyApikeys) SetDescription(v string) *GetWorkspaceResponseBodyApikeys {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBodyApikeys) SetKeyId(v string) *GetWorkspaceResponseBodyApikeys {
	s.KeyId = &v
	return s
}

func (s *GetWorkspaceResponseBodyApikeys) SetKeyName(v string) *GetWorkspaceResponseBodyApikeys {
	s.KeyName = &v
	return s
}

func (s *GetWorkspaceResponseBodyApikeys) SetKeyPrefix(v string) *GetWorkspaceResponseBodyApikeys {
	s.KeyPrefix = &v
	return s
}

func (s *GetWorkspaceResponseBodyApikeys) Validate() error {
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

type GetWorkspaceResponseBodyApikeysAuthServices struct {
	// The service ID.
	//
	// example:
	//
	// agdb-xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service type. Valid values:
	//
	// - memory
	//
	// - drama
	//
	// example:
	//
	// memory
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s GetWorkspaceResponseBodyApikeysAuthServices) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyApikeysAuthServices) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyApikeysAuthServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetWorkspaceResponseBodyApikeysAuthServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetWorkspaceResponseBodyApikeysAuthServices) SetServiceId(v string) *GetWorkspaceResponseBodyApikeysAuthServices {
	s.ServiceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyApikeysAuthServices) SetServiceType(v string) *GetWorkspaceResponseBodyApikeysAuthServices {
	s.ServiceType = &v
	return s
}

func (s *GetWorkspaceResponseBodyApikeysAuthServices) Validate() error {
	return dara.Validate(s)
}

type GetWorkspaceResponseBodyServices struct {
	// The creation time.
	//
	// example:
	//
	// 2026-03-01T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The compute resource.
	//
	// example:
	//
	// 2
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2026-06-21T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing type. Valid values:
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// - **PREPAY**: subscription.
	//
	// > - If this parameter is not specified, the default value is pay-as-you-go.
	//
	// > - In subscription billing mode, a discount is available when you purchase a duration of one year or longer. Select a billing type as needed.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The service ID.
	//
	// example:
	//
	// agdb-xxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// agdb-xxx
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service type. Valid values:
	//
	// - **memory**
	//
	// - **drama**
	//
	// example:
	//
	// memory
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service status. Valid values:
	//
	// - creating: being created.
	//
	// - active: running.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWorkspaceResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyServices) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBodyServices) GetCu() *string {
	return s.Cu
}

func (s *GetWorkspaceResponseBodyServices) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetWorkspaceResponseBodyServices) GetPayType() *string {
	return s.PayType
}

func (s *GetWorkspaceResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetWorkspaceResponseBodyServices) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetWorkspaceResponseBodyServices) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetWorkspaceResponseBodyServices) GetStatus() *string {
	return s.Status
}

func (s *GetWorkspaceResponseBodyServices) SetCreateTime(v string) *GetWorkspaceResponseBodyServices {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) SetCu(v string) *GetWorkspaceResponseBodyServices {
	s.Cu = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) SetExpireTime(v string) *GetWorkspaceResponseBodyServices {
	s.ExpireTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) SetPayType(v string) *GetWorkspaceResponseBodyServices {
	s.PayType = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) SetServiceId(v string) *GetWorkspaceResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) SetServiceName(v string) *GetWorkspaceResponseBodyServices {
	s.ServiceName = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) SetServiceType(v string) *GetWorkspaceResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) SetStatus(v string) *GetWorkspaceResponseBodyServices {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBodyServices) Validate() error {
	return dara.Validate(s)
}
