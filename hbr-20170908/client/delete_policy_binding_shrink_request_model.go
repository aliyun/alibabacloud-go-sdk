// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyBindingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceIdsShrink(v string) *DeletePolicyBindingShrinkRequest
	GetDataSourceIdsShrink() *string
	SetPolicyId(v string) *DeletePolicyBindingShrinkRequest
	GetPolicyId() *string
	SetSourceType(v string) *DeletePolicyBindingShrinkRequest
	GetSourceType() *string
}

type DeletePolicyBindingShrinkRequest struct {
	// The IDs of the data sources that you want to disassociate from the backup policy.
	DataSourceIdsShrink *string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty"`
	// The ID of the backup policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// po-000************hgp
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: ECS instance backup
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DeletePolicyBindingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyBindingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyBindingShrinkRequest) GetDataSourceIdsShrink() *string {
	return s.DataSourceIdsShrink
}

func (s *DeletePolicyBindingShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeletePolicyBindingShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DeletePolicyBindingShrinkRequest) SetDataSourceIdsShrink(v string) *DeletePolicyBindingShrinkRequest {
	s.DataSourceIdsShrink = &v
	return s
}

func (s *DeletePolicyBindingShrinkRequest) SetPolicyId(v string) *DeletePolicyBindingShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *DeletePolicyBindingShrinkRequest) SetSourceType(v string) *DeletePolicyBindingShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *DeletePolicyBindingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
