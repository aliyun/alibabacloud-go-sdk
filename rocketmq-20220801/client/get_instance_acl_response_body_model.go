// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceAclResponseBody
	GetCode() *string
	SetData(v *GetInstanceAclResponseBodyData) *GetInstanceAclResponseBody
	GetData() *GetInstanceAclResponseBodyData
	SetDynamicCode(v string) *GetInstanceAclResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetInstanceAclResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetInstanceAclResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceAclResponseBody
	GetSuccess() *bool
}

type GetInstanceAclResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetInstanceAclResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// xxx
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// xxx
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInstanceAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceAclResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceAclResponseBody) GetData() *GetInstanceAclResponseBodyData {
	return s.Data
}

func (s *GetInstanceAclResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetInstanceAclResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetInstanceAclResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceAclResponseBody) SetCode(v string) *GetInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceAclResponseBody) SetData(v *GetInstanceAclResponseBodyData) *GetInstanceAclResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceAclResponseBody) SetDynamicCode(v string) *GetInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetInstanceAclResponseBody) SetDynamicMessage(v string) *GetInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetInstanceAclResponseBody) SetHttpStatusCode(v int32) *GetInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceAclResponseBody) SetMessage(v string) *GetInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceAclResponseBody) SetRequestId(v string) *GetInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceAclResponseBody) SetSuccess(v bool) *GetInstanceAclResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceAclResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceAclResponseBodyData struct {
	// The authentication type of the instance.
	//
	// Valid values:
	//
	// 	- apache_acl: open source access control list (ACL)
	//
	// 	- default: the default account of the instance
	//
	// example:
	//
	// apache_acl
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// The type of operations that can be performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The decision result of the authorization.
	//
	// example:
	//
	// Allow
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The IP address whitelists.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the resource on which the permissions are granted.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The type of the resource on which the permissions are granted.
	//
	// example:
	//
	// Topic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The username.
	//
	// example:
	//
	// abc
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceAclResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceAclResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceAclResponseBodyData) GetAclType() *string {
	return s.AclType
}

func (s *GetInstanceAclResponseBodyData) GetActions() []*string {
	return s.Actions
}

func (s *GetInstanceAclResponseBodyData) GetDecision() *string {
	return s.Decision
}

func (s *GetInstanceAclResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceAclResponseBodyData) GetIpWhitelists() []*string {
	return s.IpWhitelists
}

func (s *GetInstanceAclResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceAclResponseBodyData) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetInstanceAclResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetInstanceAclResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *GetInstanceAclResponseBodyData) SetAclType(v string) *GetInstanceAclResponseBodyData {
	s.AclType = &v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetActions(v []*string) *GetInstanceAclResponseBodyData {
	s.Actions = v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetDecision(v string) *GetInstanceAclResponseBodyData {
	s.Decision = &v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetInstanceId(v string) *GetInstanceAclResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetIpWhitelists(v []*string) *GetInstanceAclResponseBodyData {
	s.IpWhitelists = v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetRegionId(v string) *GetInstanceAclResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetResourceName(v string) *GetInstanceAclResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetResourceType(v string) *GetInstanceAclResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *GetInstanceAclResponseBodyData) SetUsername(v string) *GetInstanceAclResponseBodyData {
	s.Username = &v
	return s
}

func (s *GetInstanceAclResponseBodyData) Validate() error {
	return dara.Validate(s)
}
