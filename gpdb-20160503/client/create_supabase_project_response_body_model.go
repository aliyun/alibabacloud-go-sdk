// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CreateSupabaseProjectResponseBody
	GetProjectId() *string
	SetRequestId(v string) *CreateSupabaseProjectResponseBody
	GetRequestId() *string
}

type CreateSupabaseProjectResponseBody struct {
	// example:
	//
	// sbp-180****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSupabaseProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSupabaseProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateSupabaseProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSupabaseProjectResponseBody) SetProjectId(v string) *CreateSupabaseProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateSupabaseProjectResponseBody) SetRequestId(v string) *CreateSupabaseProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSupabaseProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
