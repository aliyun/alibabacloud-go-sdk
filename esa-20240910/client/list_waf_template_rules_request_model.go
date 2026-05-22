// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWafTemplateRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListWafTemplateRulesRequest
	GetInstanceId() *string
	SetPhase(v string) *ListWafTemplateRulesRequest
	GetPhase() *string
	SetQueryArgs(v *ListWafTemplateRulesRequestQueryArgs) *ListWafTemplateRulesRequest
	GetQueryArgs() *ListWafTemplateRulesRequestQueryArgs
	SetSiteId(v int64) *ListWafTemplateRulesRequest
	GetSiteId() *int64
}

type ListWafTemplateRulesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// http_anti_scan
	Phase     *string                               `json:"Phase,omitempty" xml:"Phase,omitempty"`
	QueryArgs *ListWafTemplateRulesRequestQueryArgs `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty" type:"Struct"`
	SiteId    *int64                                `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListWafTemplateRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWafTemplateRulesRequest) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWafTemplateRulesRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListWafTemplateRulesRequest) GetQueryArgs() *ListWafTemplateRulesRequestQueryArgs {
	return s.QueryArgs
}

func (s *ListWafTemplateRulesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListWafTemplateRulesRequest) SetInstanceId(v string) *ListWafTemplateRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListWafTemplateRulesRequest) SetPhase(v string) *ListWafTemplateRulesRequest {
	s.Phase = &v
	return s
}

func (s *ListWafTemplateRulesRequest) SetQueryArgs(v *ListWafTemplateRulesRequestQueryArgs) *ListWafTemplateRulesRequest {
	s.QueryArgs = v
	return s
}

func (s *ListWafTemplateRulesRequest) SetSiteId(v int64) *ListWafTemplateRulesRequest {
	s.SiteId = &v
	return s
}

func (s *ListWafTemplateRulesRequest) Validate() error {
	if s.QueryArgs != nil {
		if err := s.QueryArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWafTemplateRulesRequestQueryArgs struct {
	Kinds []*string `json:"Kinds,omitempty" xml:"Kinds,omitempty" type:"Repeated"`
	// example:
	//
	// http_directory_traversal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWafTemplateRulesRequestQueryArgs) String() string {
	return dara.Prettify(s)
}

func (s ListWafTemplateRulesRequestQueryArgs) GoString() string {
	return s.String()
}

func (s *ListWafTemplateRulesRequestQueryArgs) GetKinds() []*string {
	return s.Kinds
}

func (s *ListWafTemplateRulesRequestQueryArgs) GetType() *string {
	return s.Type
}

func (s *ListWafTemplateRulesRequestQueryArgs) SetKinds(v []*string) *ListWafTemplateRulesRequestQueryArgs {
	s.Kinds = v
	return s
}

func (s *ListWafTemplateRulesRequestQueryArgs) SetType(v string) *ListWafTemplateRulesRequestQueryArgs {
	s.Type = &v
	return s
}

func (s *ListWafTemplateRulesRequestQueryArgs) Validate() error {
	return dara.Validate(s)
}
