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
	// The list of data source IDs to dissociate from the policy.
	DataSourceIdsShrink *string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// po-000************hgp
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The data source type. Valid values:
	//
	// - **UDM_ECS**: ECS instance backup.
	//
	// - **OSS**: OSS backup.
	//
	// - **NAS**: Alibaba Cloud NAS backup.
	//
	// - **COMMON_NAS**: On-premises NAS backup.
	//
	// - **ECS_FILE**: ECS File Backup Essential Edition.
	//
	// - **File**: On-premises file backup.
	//
	// - **COMMON_FILE_SYSTEM**: CPFS backup.
	//
	// - **OTS**: Tablestore backup.
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
