// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceSharedRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDataSourceSharedRuleRequest
	GetId() *int64
}

type DeleteDataSourceSharedRuleRequest struct {
	// The sharing rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22127
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataSourceSharedRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceSharedRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataSourceSharedRuleRequest) SetId(v int64) *DeleteDataSourceSharedRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataSourceSharedRuleRequest) Validate() error {
	return dara.Validate(s)
}
