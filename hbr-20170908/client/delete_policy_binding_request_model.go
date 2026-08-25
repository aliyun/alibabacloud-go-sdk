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
	// The list of data source IDs to dissociate from the policy.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
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
