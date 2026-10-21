// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIncidentInvestigationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncidentInvestigationDisplayId(v string) *GetIncidentInvestigationRequest
	GetIncidentInvestigationDisplayId() *string
	SetIncidentInvestigationId(v string) *GetIncidentInvestigationRequest
	GetIncidentInvestigationId() *string
	SetIncidentUuid(v string) *GetIncidentInvestigationRequest
	GetIncidentUuid() *string
	SetLang(v string) *GetIncidentInvestigationRequest
	GetLang() *string
	SetRegionId(v string) *GetIncidentInvestigationRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetIncidentInvestigationRequest
	GetRoleFor() *int64
}

type GetIncidentInvestigationRequest struct {
	// The display ID of the AI investigation report. You can obtain this value from the IncidentInvestigationDisplayId field returned by ListIncidentInvestigations. This parameter takes effect only when IncidentInvestigationId is not specified. If neither identifier is specified, the most recent successfully completed investigation result for the specified incident is returned.
	//
	// example:
	//
	// 755616bb5d63046c1e4f62ea872a****
	IncidentInvestigationDisplayId *string `json:"IncidentInvestigationDisplayId,omitempty" xml:"IncidentInvestigationDisplayId,omitempty"`
	// The AI investigation report ID. You can obtain this value from the IncidentInvestigationId field returned by ListIncidentInvestigations. If both this parameter and IncidentInvestigationDisplayId are specified, this parameter takes precedence.
	//
	// example:
	//
	// 951a6a0b-14c7-4332-ace6-bff21c******
	IncidentInvestigationId *string `json:"IncidentInvestigationId,omitempty" xml:"IncidentInvestigationId,omitempty"`
	// The globally unique identity of the event. If neither the report ID nor the display ID is specified, the most recent successfully completed investigation result for this event is returned via event query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Chinese mainland.
	//
	// - ap-southeast-1: Outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to the view of another member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s GetIncidentInvestigationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIncidentInvestigationRequest) GoString() string {
	return s.String()
}

func (s *GetIncidentInvestigationRequest) GetIncidentInvestigationDisplayId() *string {
	return s.IncidentInvestigationDisplayId
}

func (s *GetIncidentInvestigationRequest) GetIncidentInvestigationId() *string {
	return s.IncidentInvestigationId
}

func (s *GetIncidentInvestigationRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *GetIncidentInvestigationRequest) GetLang() *string {
	return s.Lang
}

func (s *GetIncidentInvestigationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIncidentInvestigationRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetIncidentInvestigationRequest) SetIncidentInvestigationDisplayId(v string) *GetIncidentInvestigationRequest {
	s.IncidentInvestigationDisplayId = &v
	return s
}

func (s *GetIncidentInvestigationRequest) SetIncidentInvestigationId(v string) *GetIncidentInvestigationRequest {
	s.IncidentInvestigationId = &v
	return s
}

func (s *GetIncidentInvestigationRequest) SetIncidentUuid(v string) *GetIncidentInvestigationRequest {
	s.IncidentUuid = &v
	return s
}

func (s *GetIncidentInvestigationRequest) SetLang(v string) *GetIncidentInvestigationRequest {
	s.Lang = &v
	return s
}

func (s *GetIncidentInvestigationRequest) SetRegionId(v string) *GetIncidentInvestigationRequest {
	s.RegionId = &v
	return s
}

func (s *GetIncidentInvestigationRequest) SetRoleFor(v int64) *GetIncidentInvestigationRequest {
	s.RoleFor = &v
	return s
}

func (s *GetIncidentInvestigationRequest) Validate() error {
	return dara.Validate(s)
}
