// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseOrgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateLangfuseOrgResponseBodyData) *CreateLangfuseOrgResponseBody
	GetData() *CreateLangfuseOrgResponseBodyData
	SetRequestId(v string) *CreateLangfuseOrgResponseBody
	GetRequestId() *string
}

type CreateLangfuseOrgResponseBody struct {
	// The response data.
	Data *CreateLangfuseOrgResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLangfuseOrgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgResponseBody) GetData() *CreateLangfuseOrgResponseBodyData {
	return s.Data
}

func (s *CreateLangfuseOrgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLangfuseOrgResponseBody) SetData(v *CreateLangfuseOrgResponseBodyData) *CreateLangfuseOrgResponseBody {
	s.Data = v
	return s
}

func (s *CreateLangfuseOrgResponseBody) SetRequestId(v string) *CreateLangfuseOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLangfuseOrgResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLangfuseOrgResponseBodyData struct {
	// The time when the Langfuse organization was created.
	//
	// example:
	//
	// 2026-06-25T09:28:30.949Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The name of the Langfuse organization.
	//
	// example:
	//
	// org_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Langfuse organization ID.
	//
	// example:
	//
	// cmrbhzx930005jw****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s CreateLangfuseOrgResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateLangfuseOrgResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateLangfuseOrgResponseBodyData) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateLangfuseOrgResponseBodyData) SetCreatedAt(v string) *CreateLangfuseOrgResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateLangfuseOrgResponseBodyData) SetName(v string) *CreateLangfuseOrgResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateLangfuseOrgResponseBodyData) SetOrganizationId(v string) *CreateLangfuseOrgResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *CreateLangfuseOrgResponseBodyData) Validate() error {
	return dara.Validate(s)
}
