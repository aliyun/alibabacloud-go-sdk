// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDataQualityRuleRequest
	GetId() *int64
	SetProjectId(v int64) *DeleteDataQualityRuleRequest
	GetProjectId() *int64
}

type DeleteDataQualityRuleRequest struct {
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19715
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 17302
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDataQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataQualityRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DeleteDataQualityRuleRequest) SetId(v int64) *DeleteDataQualityRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataQualityRuleRequest) SetProjectId(v int64) *DeleteDataQualityRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDataQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
