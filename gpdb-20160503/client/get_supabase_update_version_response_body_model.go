// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseUpdateVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatestVersion(v string) *GetSupabaseUpdateVersionResponseBody
	GetLatestVersion() *string
	SetProjectId(v string) *GetSupabaseUpdateVersionResponseBody
	GetProjectId() *string
	SetRequestId(v string) *GetSupabaseUpdateVersionResponseBody
	GetRequestId() *string
	SetStableVersion(v string) *GetSupabaseUpdateVersionResponseBody
	GetStableVersion() *string
}

type GetSupabaseUpdateVersionResponseBody struct {
	// The latest upgradable version.
	//
	// example:
	//
	// 20240731
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The ID of the Supabase project.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The recommended stable version for upgrade.
	//
	// example:
	//
	// 20240630
	StableVersion *string `json:"StableVersion,omitempty" xml:"StableVersion,omitempty"`
}

func (s GetSupabaseUpdateVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseUpdateVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupabaseUpdateVersionResponseBody) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetSupabaseUpdateVersionResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSupabaseUpdateVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupabaseUpdateVersionResponseBody) GetStableVersion() *string {
	return s.StableVersion
}

func (s *GetSupabaseUpdateVersionResponseBody) SetLatestVersion(v string) *GetSupabaseUpdateVersionResponseBody {
	s.LatestVersion = &v
	return s
}

func (s *GetSupabaseUpdateVersionResponseBody) SetProjectId(v string) *GetSupabaseUpdateVersionResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetSupabaseUpdateVersionResponseBody) SetRequestId(v string) *GetSupabaseUpdateVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupabaseUpdateVersionResponseBody) SetStableVersion(v string) *GetSupabaseUpdateVersionResponseBody {
	s.StableVersion = &v
	return s
}

func (s *GetSupabaseUpdateVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
