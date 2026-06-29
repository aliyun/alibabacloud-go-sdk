// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateOrganizationResponseBody
	GetCode() *string
	SetData(v *UpdateOrganizationResponseBodyData) *UpdateOrganizationResponseBody
	GetData() *UpdateOrganizationResponseBodyData
	SetHttpStatusCode(v int32) *UpdateOrganizationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateOrganizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOrganizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOrganizationResponseBody
	GetSuccess() *bool
}

type UpdateOrganizationResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *UpdateOrganizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// None
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC0B0D42-A88C-55F5-AF48-498B8207F6BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOrganizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateOrganizationResponseBody) GetData() *UpdateOrganizationResponseBodyData {
	return s.Data
}

func (s *UpdateOrganizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateOrganizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOrganizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOrganizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOrganizationResponseBody) SetCode(v string) *UpdateOrganizationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOrganizationResponseBody) SetData(v *UpdateOrganizationResponseBodyData) *UpdateOrganizationResponseBody {
	s.Data = v
	return s
}

func (s *UpdateOrganizationResponseBody) SetHttpStatusCode(v int32) *UpdateOrganizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateOrganizationResponseBody) SetMessage(v string) *UpdateOrganizationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOrganizationResponseBody) SetRequestId(v string) *UpdateOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOrganizationResponseBody) SetSuccess(v bool) *UpdateOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOrganizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateOrganizationResponseBodyData struct {
	// The organization description.
	//
	// example:
	//
	// 组织描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-12-11T02:19:27Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-07-08 13:59:11
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The organization name.
	//
	// example:
	//
	// 组织名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The product namespace ID.
	//
	// example:
	//
	// namespace-1
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// org_123456789
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The business account identifier of the organization owner. The value is aliyunUid for the ALIYUN type and userIdentifier for the SSO type.
	//
	// example:
	//
	// 1269388320468566
	OwnerBizAccountId *string `json:"OwnerBizAccountId,omitempty" xml:"OwnerBizAccountId,omitempty"`
	// The UAC account ID of the organization owner.
	//
	// example:
	//
	// acc_123456789
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The status. Valid values:
	//
	// - ACTIVE
	//
	// - FROZEN
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateOrganizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateOrganizationResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateOrganizationResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateOrganizationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateOrganizationResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateOrganizationResponseBodyData) GetOrgId() *string {
	return s.OrgId
}

func (s *UpdateOrganizationResponseBodyData) GetOwnerBizAccountId() *string {
	return s.OwnerBizAccountId
}

func (s *UpdateOrganizationResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *UpdateOrganizationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateOrganizationResponseBodyData) SetDescription(v string) *UpdateOrganizationResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetGmtCreate(v string) *UpdateOrganizationResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetGmtModified(v string) *UpdateOrganizationResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetName(v string) *UpdateOrganizationResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetNamespaceId(v string) *UpdateOrganizationResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetOrgId(v string) *UpdateOrganizationResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetOwnerBizAccountId(v string) *UpdateOrganizationResponseBodyData {
	s.OwnerBizAccountId = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetOwnerId(v string) *UpdateOrganizationResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) SetStatus(v string) *UpdateOrganizationResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateOrganizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
