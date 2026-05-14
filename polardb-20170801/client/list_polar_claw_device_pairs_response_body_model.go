// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolarClawDevicePairsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListPolarClawDevicePairsResponseBody
	GetApplicationId() *string
	SetCode(v int32) *ListPolarClawDevicePairsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListPolarClawDevicePairsResponseBody
	GetMessage() *string
	SetPaired(v []*ListPolarClawDevicePairsResponseBodyPaired) *ListPolarClawDevicePairsResponseBody
	GetPaired() []*ListPolarClawDevicePairsResponseBodyPaired
	SetPending(v []*ListPolarClawDevicePairsResponseBodyPending) *ListPolarClawDevicePairsResponseBody
	GetPending() []*ListPolarClawDevicePairsResponseBodyPending
	SetRequestId(v string) *ListPolarClawDevicePairsResponseBody
	GetRequestId() *string
}

type ListPolarClawDevicePairsResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Paired  []*ListPolarClawDevicePairsResponseBodyPaired  `json:"Paired,omitempty" xml:"Paired,omitempty" type:"Repeated"`
	Pending []*ListPolarClawDevicePairsResponseBodyPending `json:"Pending,omitempty" xml:"Pending,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPolarClawDevicePairsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawDevicePairsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolarClawDevicePairsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPolarClawDevicePairsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPolarClawDevicePairsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPolarClawDevicePairsResponseBody) GetPaired() []*ListPolarClawDevicePairsResponseBodyPaired {
	return s.Paired
}

func (s *ListPolarClawDevicePairsResponseBody) GetPending() []*ListPolarClawDevicePairsResponseBodyPending {
	return s.Pending
}

func (s *ListPolarClawDevicePairsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolarClawDevicePairsResponseBody) SetApplicationId(v string) *ListPolarClawDevicePairsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBody) SetCode(v int32) *ListPolarClawDevicePairsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBody) SetMessage(v string) *ListPolarClawDevicePairsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBody) SetPaired(v []*ListPolarClawDevicePairsResponseBodyPaired) *ListPolarClawDevicePairsResponseBody {
	s.Paired = v
	return s
}

func (s *ListPolarClawDevicePairsResponseBody) SetPending(v []*ListPolarClawDevicePairsResponseBodyPending) *ListPolarClawDevicePairsResponseBody {
	s.Pending = v
	return s
}

func (s *ListPolarClawDevicePairsResponseBody) SetRequestId(v string) *ListPolarClawDevicePairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBody) Validate() error {
	if s.Paired != nil {
		for _, item := range s.Paired {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pending != nil {
		for _, item := range s.Pending {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolarClawDevicePairsResponseBodyPaired struct {
	// example:
	//
	// cli
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// cli
	ClientMode *string `json:"ClientMode,omitempty" xml:"ClientMode,omitempty"`
	// example:
	//
	// 1778659807727
	CreatedAtMs *int64 `json:"CreatedAtMs,omitempty" xml:"CreatedAtMs,omitempty"`
	// example:
	//
	// empty
	DeviceFamily *string `json:"DeviceFamily,omitempty" xml:"DeviceFamily,omitempty"`
	// example:
	//
	// f92620d6bea04f65d48cf603c57b367c97e837c1ab9f6d78f741f477e99d857c
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// empty
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1778659807727
	LastSeenAtMs *int64 `json:"LastSeenAtMs,omitempty" xml:"LastSeenAtMs,omitempty"`
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// operator
	Role   *string   `json:"Role,omitempty" xml:"Role,omitempty"`
	Scopes []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
}

func (s ListPolarClawDevicePairsResponseBodyPaired) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawDevicePairsResponseBodyPaired) GoString() string {
	return s.String()
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetClientId() *string {
	return s.ClientId
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetClientMode() *string {
	return s.ClientMode
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetCreatedAtMs() *int64 {
	return s.CreatedAtMs
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetDeviceFamily() *string {
	return s.DeviceFamily
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetLastSeenAtMs() *int64 {
	return s.LastSeenAtMs
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetPlatform() *string {
	return s.Platform
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetRole() *string {
	return s.Role
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) GetScopes() []*string {
	return s.Scopes
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetClientId(v string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.ClientId = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetClientMode(v string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.ClientMode = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetCreatedAtMs(v int64) *ListPolarClawDevicePairsResponseBodyPaired {
	s.CreatedAtMs = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetDeviceFamily(v string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.DeviceFamily = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetDeviceId(v string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.DeviceId = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetDisplayName(v string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.DisplayName = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetLastSeenAtMs(v int64) *ListPolarClawDevicePairsResponseBodyPaired {
	s.LastSeenAtMs = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetPlatform(v string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.Platform = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetRole(v string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.Role = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) SetScopes(v []*string) *ListPolarClawDevicePairsResponseBodyPaired {
	s.Scopes = v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPaired) Validate() error {
	return dara.Validate(s)
}

type ListPolarClawDevicePairsResponseBodyPending struct {
	// example:
	//
	// cli
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// cli
	ClientMode *string `json:"ClientMode,omitempty" xml:"ClientMode,omitempty"`
	// example:
	//
	// server
	DeviceFamily *string `json:"DeviceFamily,omitempty" xml:"DeviceFamily,omitempty"`
	// example:
	//
	// device-784x37k0vko734fk
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// false
	IsRepair *bool `json:"IsRepair,omitempty" xml:"IsRepair,omitempty"`
	// example:
	//
	// test-1778647932986
	PairRequestId *string `json:"PairRequestId,omitempty" xml:"PairRequestId,omitempty"`
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// test-pubkey-base64url
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// example:
	//
	// 127.0.0.1
	RemoteIp *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	// example:
	//
	// operator
	Role   *string   `json:"Role,omitempty" xml:"Role,omitempty"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	Scopes []*string `json:"Scopes,omitempty" xml:"Scopes,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Silent *bool `json:"Silent,omitempty" xml:"Silent,omitempty"`
	// example:
	//
	// 1778660347550
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s ListPolarClawDevicePairsResponseBodyPending) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawDevicePairsResponseBodyPending) GoString() string {
	return s.String()
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetClientId() *string {
	return s.ClientId
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetClientMode() *string {
	return s.ClientMode
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetDeviceFamily() *string {
	return s.DeviceFamily
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetIsRepair() *bool {
	return s.IsRepair
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetPairRequestId() *string {
	return s.PairRequestId
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetPlatform() *string {
	return s.Platform
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetPublicKey() *string {
	return s.PublicKey
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetRemoteIp() *string {
	return s.RemoteIp
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetRole() *string {
	return s.Role
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetRoles() []*string {
	return s.Roles
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetScopes() []*string {
	return s.Scopes
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetSilent() *bool {
	return s.Silent
}

func (s *ListPolarClawDevicePairsResponseBodyPending) GetTs() *int64 {
	return s.Ts
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetClientId(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.ClientId = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetClientMode(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.ClientMode = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetDeviceFamily(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.DeviceFamily = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetDeviceId(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.DeviceId = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetDisplayName(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.DisplayName = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetIsRepair(v bool) *ListPolarClawDevicePairsResponseBodyPending {
	s.IsRepair = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetPairRequestId(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.PairRequestId = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetPlatform(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.Platform = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetPublicKey(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.PublicKey = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetRemoteIp(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.RemoteIp = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetRole(v string) *ListPolarClawDevicePairsResponseBodyPending {
	s.Role = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetRoles(v []*string) *ListPolarClawDevicePairsResponseBodyPending {
	s.Roles = v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetScopes(v []*string) *ListPolarClawDevicePairsResponseBodyPending {
	s.Scopes = v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetSilent(v bool) *ListPolarClawDevicePairsResponseBodyPending {
	s.Silent = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) SetTs(v int64) *ListPolarClawDevicePairsResponseBodyPending {
	s.Ts = &v
	return s
}

func (s *ListPolarClawDevicePairsResponseBodyPending) Validate() error {
	return dara.Validate(s)
}
