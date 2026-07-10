// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseOrgMembershipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteLangfuseOrgMembershipResponseBodyData) *DeleteLangfuseOrgMembershipResponseBody
	GetData() *DeleteLangfuseOrgMembershipResponseBodyData
	SetRequestId(v string) *DeleteLangfuseOrgMembershipResponseBody
	GetRequestId() *string
}

type DeleteLangfuseOrgMembershipResponseBody struct {
	// The response result.
	Data *DeleteLangfuseOrgMembershipResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLangfuseOrgMembershipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgMembershipResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgMembershipResponseBody) GetData() *DeleteLangfuseOrgMembershipResponseBodyData {
	return s.Data
}

func (s *DeleteLangfuseOrgMembershipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLangfuseOrgMembershipResponseBody) SetData(v *DeleteLangfuseOrgMembershipResponseBodyData) *DeleteLangfuseOrgMembershipResponseBody {
	s.Data = v
	return s
}

func (s *DeleteLangfuseOrgMembershipResponseBody) SetRequestId(v string) *DeleteLangfuseOrgMembershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLangfuseOrgMembershipResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLangfuseOrgMembershipResponseBodyData struct {
	// The email address of the user.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s DeleteLangfuseOrgMembershipResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgMembershipResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgMembershipResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *DeleteLangfuseOrgMembershipResponseBodyData) SetEmail(v string) *DeleteLangfuseOrgMembershipResponseBodyData {
	s.Email = &v
	return s
}

func (s *DeleteLangfuseOrgMembershipResponseBodyData) Validate() error {
	return dara.Validate(s)
}
