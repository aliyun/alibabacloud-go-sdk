// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseUserAllowOrgCreationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyLangfuseUserAllowOrgCreationResponseBodyData) *ModifyLangfuseUserAllowOrgCreationResponseBody
	GetData() *ModifyLangfuseUserAllowOrgCreationResponseBodyData
	SetRequestId(v string) *ModifyLangfuseUserAllowOrgCreationResponseBody
	GetRequestId() *string
}

type ModifyLangfuseUserAllowOrgCreationResponseBody struct {
	// The returned result.
	Data *ModifyLangfuseUserAllowOrgCreationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLangfuseUserAllowOrgCreationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseUserAllowOrgCreationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBody) GetData() *ModifyLangfuseUserAllowOrgCreationResponseBodyData {
	return s.Data
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBody) SetData(v *ModifyLangfuseUserAllowOrgCreationResponseBodyData) *ModifyLangfuseUserAllowOrgCreationResponseBody {
	s.Data = v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBody) SetRequestId(v string) *ModifyLangfuseUserAllowOrgCreationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyLangfuseUserAllowOrgCreationResponseBodyData struct {
	// Specifies whether the user is allowed to create organizations.
	//
	// example:
	//
	// false
	AllowOrgCreation *bool `json:"AllowOrgCreation,omitempty" xml:"AllowOrgCreation,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// alice@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s ModifyLangfuseUserAllowOrgCreationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseUserAllowOrgCreationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBodyData) GetAllowOrgCreation() *bool {
	return s.AllowOrgCreation
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBodyData) SetAllowOrgCreation(v bool) *ModifyLangfuseUserAllowOrgCreationResponseBodyData {
	s.AllowOrgCreation = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBodyData) SetEmail(v string) *ModifyLangfuseUserAllowOrgCreationResponseBodyData {
	s.Email = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
