// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectDashboardAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDashboardPassword(v string) *GetSupabaseProjectDashboardAccountResponseBody
	GetDashboardPassword() *string
	SetDashboardUsername(v string) *GetSupabaseProjectDashboardAccountResponseBody
	GetDashboardUsername() *string
	SetProjectId(v string) *GetSupabaseProjectDashboardAccountResponseBody
	GetProjectId() *string
	SetProjectName(v string) *GetSupabaseProjectDashboardAccountResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetSupabaseProjectDashboardAccountResponseBody
	GetRequestId() *string
}

type GetSupabaseProjectDashboardAccountResponseBody struct {
	// The username for accessing the project\\"s dashboard.
	//
	// example:
	//
	// xxpassword
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// The password associated with the dashboard username.
	//
	// example:
	//
	// xxuser
	DashboardUsername *string `json:"DashboardUsername,omitempty" xml:"DashboardUsername,omitempty"`
	// The ID of the Supabase project.
	//
	// example:
	//
	// sbp-twmoe9bakow
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the Supabase project.
	//
	// example:
	//
	// supabase_prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSupabaseProjectDashboardAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectDashboardAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) GetDashboardUsername() *string {
	return s.DashboardUsername
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) SetDashboardPassword(v string) *GetSupabaseProjectDashboardAccountResponseBody {
	s.DashboardPassword = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) SetDashboardUsername(v string) *GetSupabaseProjectDashboardAccountResponseBody {
	s.DashboardUsername = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) SetProjectId(v string) *GetSupabaseProjectDashboardAccountResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) SetProjectName(v string) *GetSupabaseProjectDashboardAccountResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) SetRequestId(v string) *GetSupabaseProjectDashboardAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupabaseProjectDashboardAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
