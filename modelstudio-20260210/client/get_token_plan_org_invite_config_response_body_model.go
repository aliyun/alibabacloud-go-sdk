// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanOrgInviteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenPlanOrgInviteConfigResponseBody
	GetCode() *string
	SetData(v *GetTokenPlanOrgInviteConfigResponseBodyData) *GetTokenPlanOrgInviteConfigResponseBody
	GetData() *GetTokenPlanOrgInviteConfigResponseBodyData
	SetMessage(v string) *GetTokenPlanOrgInviteConfigResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetTokenPlanOrgInviteConfigResponseBody
	GetSuccess() *bool
}

type GetTokenPlanOrgInviteConfigResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data result of the current category statistics.
	Data *GetTokenPlanOrgInviteConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTokenPlanOrgInviteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanOrgInviteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) GetData() *GetTokenPlanOrgInviteConfigResponseBodyData {
	return s.Data
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) SetCode(v string) *GetTokenPlanOrgInviteConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) SetData(v *GetTokenPlanOrgInviteConfigResponseBodyData) *GetTokenPlanOrgInviteConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) SetMessage(v string) *GetTokenPlanOrgInviteConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) SetSuccess(v bool) *GetTokenPlanOrgInviteConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenPlanOrgInviteConfigResponseBodyData struct {
	// The default organization role ID to assign. Valid values:
	//
	// - SYSTEM_ROLE_ORG_ADMIN
	//
	// - SYSTEM_ROLE_ORG_MEMBER
	//
	// example:
	//
	// ORG_MEMBER
	DefaultRoleId *string `json:"DefaultRoleId,omitempty" xml:"DefaultRoleId,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// org_123456789
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The default seat allocation strategy. Valid values:
	//
	// - HIGH_TO_LOW
	//
	// - LOW_TO_HIGH
	//
	// - NONE
	//
	// example:
	//
	// NONE
	SeatAssignStrategy *string `json:"SeatAssignStrategy,omitempty" xml:"SeatAssignStrategy,omitempty"`
}

func (s GetTokenPlanOrgInviteConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanOrgInviteConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenPlanOrgInviteConfigResponseBodyData) GetDefaultRoleId() *string {
	return s.DefaultRoleId
}

func (s *GetTokenPlanOrgInviteConfigResponseBodyData) GetOrgId() *string {
	return s.OrgId
}

func (s *GetTokenPlanOrgInviteConfigResponseBodyData) GetSeatAssignStrategy() *string {
	return s.SeatAssignStrategy
}

func (s *GetTokenPlanOrgInviteConfigResponseBodyData) SetDefaultRoleId(v string) *GetTokenPlanOrgInviteConfigResponseBodyData {
	s.DefaultRoleId = &v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponseBodyData) SetOrgId(v string) *GetTokenPlanOrgInviteConfigResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponseBodyData) SetSeatAssignStrategy(v string) *GetTokenPlanOrgInviteConfigResponseBodyData {
	s.SeatAssignStrategy = &v
	return s
}

func (s *GetTokenPlanOrgInviteConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
