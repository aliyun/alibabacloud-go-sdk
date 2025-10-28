// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHDMInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddHDMInstanceResponseBody
	GetCode() *string
	SetData(v *AddHDMInstanceResponseBodyData) *AddHDMInstanceResponseBody
	GetData() *AddHDMInstanceResponseBodyData
	SetMessage(v string) *AddHDMInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddHDMInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AddHDMInstanceResponseBody
	GetSuccess() *string
	SetSynchro(v string) *AddHDMInstanceResponseBody
	GetSynchro() *string
}

type AddHDMInstanceResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information, including the error codes and the number of entries that are returned.
	Data *AddHDMInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Synchro *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s AddHDMInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddHDMInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddHDMInstanceResponseBody) GetData() *AddHDMInstanceResponseBodyData {
	return s.Data
}

func (s *AddHDMInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddHDMInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddHDMInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AddHDMInstanceResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *AddHDMInstanceResponseBody) SetCode(v string) *AddHDMInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetData(v *AddHDMInstanceResponseBodyData) *AddHDMInstanceResponseBody {
	s.Data = v
	return s
}

func (s *AddHDMInstanceResponseBody) SetMessage(v string) *AddHDMInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetRequestId(v string) *AddHDMInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetSuccess(v string) *AddHDMInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *AddHDMInstanceResponseBody) SetSynchro(v string) *AddHDMInstanceResponseBody {
	s.Synchro = &v
	return s
}

func (s *AddHDMInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddHDMInstanceResponseBodyData struct {
	// The user ID of the caller.
	//
	// example:
	//
	// 31063db679****
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// InvalidRequestURL
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// rm-de21209****.mysql.rds.aliyuncs.com
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the instance owner.
	//
	// example:
	//
	// 325352345
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number of the instance that you want to access.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The role of the current API caller.
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// L0EPfLS****=SCE00000*****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// tokenID
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The unique identifier of the instance.
	//
	// example:
	//
	// hdm_3063db6792965c080a4bcb6e6304****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-m5e666n89m2bx8jar****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AddHDMInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddHDMInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddHDMInstanceResponseBodyData) GetCallerUid() *string {
	return s.CallerUid
}

func (s *AddHDMInstanceResponseBodyData) GetCode() *int32 {
	return s.Code
}

func (s *AddHDMInstanceResponseBodyData) GetError() *string {
	return s.Error
}

func (s *AddHDMInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddHDMInstanceResponseBodyData) GetIp() *string {
	return s.Ip
}

func (s *AddHDMInstanceResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AddHDMInstanceResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *AddHDMInstanceResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *AddHDMInstanceResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *AddHDMInstanceResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *AddHDMInstanceResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *AddHDMInstanceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *AddHDMInstanceResponseBodyData) SetCallerUid(v string) *AddHDMInstanceResponseBodyData {
	s.CallerUid = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetCode(v int32) *AddHDMInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetError(v string) *AddHDMInstanceResponseBodyData {
	s.Error = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetInstanceId(v string) *AddHDMInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetIp(v string) *AddHDMInstanceResponseBodyData {
	s.Ip = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetOwnerId(v string) *AddHDMInstanceResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetPort(v int32) *AddHDMInstanceResponseBodyData {
	s.Port = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetRole(v string) *AddHDMInstanceResponseBodyData {
	s.Role = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetTenantId(v string) *AddHDMInstanceResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetToken(v string) *AddHDMInstanceResponseBodyData {
	s.Token = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetUuid(v string) *AddHDMInstanceResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) SetVpcId(v string) *AddHDMInstanceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *AddHDMInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
