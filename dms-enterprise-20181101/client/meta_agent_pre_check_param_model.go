// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaAgentPreCheckParam interface {
	dara.Model
	String() string
	GoString() string
	SetDbIds(v string) *MetaAgentPreCheckParam
	GetDbIds() *string
	SetDescription(v string) *MetaAgentPreCheckParam
	GetDescription() *string
	SetInstanceIds(v string) *MetaAgentPreCheckParam
	GetInstanceIds() *string
	SetSupplement(v string) *MetaAgentPreCheckParam
	GetSupplement() *string
	SetTableNames(v string) *MetaAgentPreCheckParam
	GetTableNames() *string
}

type MetaAgentPreCheckParam struct {
	DbIds       *string `json:"DbIds,omitempty" xml:"DbIds,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Supplement  *string `json:"Supplement,omitempty" xml:"Supplement,omitempty"`
	TableNames  *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
}

func (s MetaAgentPreCheckParam) String() string {
	return dara.Prettify(s)
}

func (s MetaAgentPreCheckParam) GoString() string {
	return s.String()
}

func (s *MetaAgentPreCheckParam) GetDbIds() *string {
	return s.DbIds
}

func (s *MetaAgentPreCheckParam) GetDescription() *string {
	return s.Description
}

func (s *MetaAgentPreCheckParam) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *MetaAgentPreCheckParam) GetSupplement() *string {
	return s.Supplement
}

func (s *MetaAgentPreCheckParam) GetTableNames() *string {
	return s.TableNames
}

func (s *MetaAgentPreCheckParam) SetDbIds(v string) *MetaAgentPreCheckParam {
	s.DbIds = &v
	return s
}

func (s *MetaAgentPreCheckParam) SetDescription(v string) *MetaAgentPreCheckParam {
	s.Description = &v
	return s
}

func (s *MetaAgentPreCheckParam) SetInstanceIds(v string) *MetaAgentPreCheckParam {
	s.InstanceIds = &v
	return s
}

func (s *MetaAgentPreCheckParam) SetSupplement(v string) *MetaAgentPreCheckParam {
	s.Supplement = &v
	return s
}

func (s *MetaAgentPreCheckParam) SetTableNames(v string) *MetaAgentPreCheckParam {
	s.TableNames = &v
	return s
}

func (s *MetaAgentPreCheckParam) Validate() error {
	return dara.Validate(s)
}
