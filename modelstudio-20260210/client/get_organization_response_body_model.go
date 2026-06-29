// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOrganizationResponseBody
	GetCode() *string
	SetData(v *GetOrganizationResponseBodyData) *GetOrganizationResponseBody
	GetData() *GetOrganizationResponseBodyData
	SetHttpStatusCode(v int32) *GetOrganizationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetOrganizationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOrganizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrganizationResponseBody
	GetSuccess() *bool
}

type GetOrganizationResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetOrganizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 0CB5868D-C6E5-59A6-A20A-C39EB2E75BDE
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

func (s GetOrganizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOrganizationResponseBody) GetData() *GetOrganizationResponseBodyData {
	return s.Data
}

func (s *GetOrganizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetOrganizationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOrganizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrganizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrganizationResponseBody) SetCode(v string) *GetOrganizationResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrganizationResponseBody) SetData(v *GetOrganizationResponseBodyData) *GetOrganizationResponseBody {
	s.Data = v
	return s
}

func (s *GetOrganizationResponseBody) SetHttpStatusCode(v int32) *GetOrganizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOrganizationResponseBody) SetMessage(v string) *GetOrganizationResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrganizationResponseBody) SetRequestId(v string) *GetOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrganizationResponseBody) SetSuccess(v bool) *GetOrganizationResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrganizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrganizationResponseBodyData struct {
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
	// 2025-11-20T02:26:35Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-11-13T02:11:56Z
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
	// 1543686331379464
	OwnerBizAccountId *string `json:"OwnerBizAccountId,omitempty" xml:"OwnerBizAccountId,omitempty"`
	// The UAC account ID of the organization owner.
	//
	// example:
	//
	// acc_123456789
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The status. Valid values: ACTIVE and FROZEN.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetOrganizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrganizationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetOrganizationResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetOrganizationResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetOrganizationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetOrganizationResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetOrganizationResponseBodyData) GetOrgId() *string {
	return s.OrgId
}

func (s *GetOrganizationResponseBodyData) GetOwnerBizAccountId() *string {
	return s.OwnerBizAccountId
}

func (s *GetOrganizationResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetOrganizationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetOrganizationResponseBodyData) SetDescription(v string) *GetOrganizationResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetGmtCreate(v string) *GetOrganizationResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetGmtModified(v string) *GetOrganizationResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetName(v string) *GetOrganizationResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetNamespaceId(v string) *GetOrganizationResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetOrgId(v string) *GetOrganizationResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetOwnerBizAccountId(v string) *GetOrganizationResponseBodyData {
	s.OwnerBizAccountId = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetOwnerId(v string) *GetOrganizationResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetOrganizationResponseBodyData) SetStatus(v string) *GetOrganizationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetOrganizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
