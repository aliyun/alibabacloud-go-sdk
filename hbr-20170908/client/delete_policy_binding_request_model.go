// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceIds(v []*string) *DeletePolicyBindingRequest
	GetDataSourceIds() []*string
	SetPolicyId(v string) *DeletePolicyBindingRequest
	GetPolicyId() *string
	SetSourceType(v string) *DeletePolicyBindingRequest
	GetSourceType() *string
}

type DeletePolicyBindingRequest struct {
	// The IDs of the data sources that you want to disassociate from the backup policy.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
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

func (s DeletePolicyBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyBindingRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyBindingRequest) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *DeletePolicyBindingRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeletePolicyBindingRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DeletePolicyBindingRequest) SetDataSourceIds(v []*string) *DeletePolicyBindingRequest {
	s.DataSourceIds = v
	return s
}

func (s *DeletePolicyBindingRequest) SetPolicyId(v string) *DeletePolicyBindingRequest {
	s.PolicyId = &v
	return s
}

func (s *DeletePolicyBindingRequest) SetSourceType(v string) *DeletePolicyBindingRequest {
	s.SourceType = &v
	return s
}

func (s *DeletePolicyBindingRequest) Validate() error {
	return dara.Validate(s)
}
