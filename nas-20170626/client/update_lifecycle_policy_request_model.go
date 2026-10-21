// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLifecyclePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteRules(v []*UpdateLifecyclePolicyRequestDeleteRules) *UpdateLifecyclePolicyRequest
	GetDeleteRules() []*UpdateLifecyclePolicyRequestDeleteRules
	SetDescription(v string) *UpdateLifecyclePolicyRequest
	GetDescription() *string
	SetFileSystemId(v string) *UpdateLifecyclePolicyRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *UpdateLifecyclePolicyRequest
	GetLifecyclePolicyId() *string
	SetPaths(v []*string) *UpdateLifecyclePolicyRequest
	GetPaths() []*string
	SetRetrieveRules(v []*UpdateLifecyclePolicyRequestRetrieveRules) *UpdateLifecyclePolicyRequest
	GetRetrieveRules() []*UpdateLifecyclePolicyRequestRetrieveRules
	SetStorageType(v string) *UpdateLifecyclePolicyRequest
	GetStorageType() *string
	SetTransitRules(v []*UpdateLifecyclePolicyRequestTransitRules) *UpdateLifecyclePolicyRequest
	GetTransitRules() []*UpdateLifecyclePolicyRequestTransitRules
}

type UpdateLifecyclePolicyRequest struct {
	// The expiration and deletion rules for file data.
	DeleteRules []*UpdateLifecyclePolicyRequestDeleteRules `json:"DeleteRules,omitempty" xml:"DeleteRules,omitempty" type:"Repeated"`
	// The description of the lifecycle policy.
	//
	// Format:
	//
	// The description must be 3 to 64 characters in length. It must start with a letter and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// > Supported only for CPFS for Lingjun.
	//
	// example:
	//
	// Lifecycle policy description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system. The ID starts with bmcpfs-, such as bmcpfs-290w65p03ok64ya****. To use CPFS for Lingjun, you must first submit a ticket to apply for the service whitelist and then create the file system in the NAS console.
	//
	// > Supported only when LifecyclePolicyType is set to OnDemand in the lifecycle management policy of a CPFS for Lingjun file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64y*****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the lifecycle policy.
	//
	// > Required for CPFS for Lingjun file systems.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsp-bp1234567890ab****
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	// The absolute path of the directory associated with the lifecycle management policy. This parameter is required when FileSystemId starts with bmcpfs-.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The data retrieval rules for files. You can configure up to one rule.
	//
	// > Supported only for CPFS for Lingjun file systems.
	RetrieveRules []*UpdateLifecyclePolicyRequestRetrieveRules `json:"RetrieveRules,omitempty" xml:"RetrieveRules,omitempty" type:"Repeated"`
	// The tiered storage type. Valid values:
	//
	// - InfrequentAccess: IA storage class. This is the default value.
	//
	// - Archive: Archive storage class.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The data transition rules for files. You can configure up to one rule.
	//
	// > Supported only when LifecyclePolicyType is set to Auto for a CPFS for Lingjun file system.
	TransitRules []*UpdateLifecyclePolicyRequestTransitRules `json:"TransitRules,omitempty" xml:"TransitRules,omitempty" type:"Repeated"`
}

func (s UpdateLifecyclePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyRequest) GetDeleteRules() []*UpdateLifecyclePolicyRequestDeleteRules {
	return s.DeleteRules
}

func (s *UpdateLifecyclePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLifecyclePolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *UpdateLifecyclePolicyRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *UpdateLifecyclePolicyRequest) GetPaths() []*string {
	return s.Paths
}

func (s *UpdateLifecyclePolicyRequest) GetRetrieveRules() []*UpdateLifecyclePolicyRequestRetrieveRules {
	return s.RetrieveRules
}

func (s *UpdateLifecyclePolicyRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *UpdateLifecyclePolicyRequest) GetTransitRules() []*UpdateLifecyclePolicyRequestTransitRules {
	return s.TransitRules
}

func (s *UpdateLifecyclePolicyRequest) SetDeleteRules(v []*UpdateLifecyclePolicyRequestDeleteRules) *UpdateLifecyclePolicyRequest {
	s.DeleteRules = v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetDescription(v string) *UpdateLifecyclePolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetFileSystemId(v string) *UpdateLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetLifecyclePolicyId(v string) *UpdateLifecyclePolicyRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetPaths(v []*string) *UpdateLifecyclePolicyRequest {
	s.Paths = v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetRetrieveRules(v []*UpdateLifecyclePolicyRequestRetrieveRules) *UpdateLifecyclePolicyRequest {
	s.RetrieveRules = v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetStorageType(v string) *UpdateLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

func (s *UpdateLifecyclePolicyRequest) SetTransitRules(v []*UpdateLifecyclePolicyRequestTransitRules) *UpdateLifecyclePolicyRequest {
	s.TransitRules = v
	return s
}

func (s *UpdateLifecyclePolicyRequest) Validate() error {
	if s.DeleteRules != nil {
		for _, item := range s.DeleteRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RetrieveRules != nil {
		for _, item := range s.RetrieveRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TransitRules != nil {
		for _, item := range s.TransitRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateLifecyclePolicyRequestDeleteRules struct {
	// The attribute of the rule.
	//
	// Valid values:
	//
	// - Atime: the last access time of the file.
	//
	// example:
	//
	// Atime
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The threshold of the rule.
	//
	// Valid values:
	//
	// - If Attribute is set to Atime, the value indicates the number of days that a file has not been accessed. Valid values: 1 to 365.
	//
	// example:
	//
	// 4
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLifecyclePolicyRequestDeleteRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyRequestDeleteRules) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyRequestDeleteRules) GetAttribute() *string {
	return s.Attribute
}

func (s *UpdateLifecyclePolicyRequestDeleteRules) GetThreshold() *string {
	return s.Threshold
}

func (s *UpdateLifecyclePolicyRequestDeleteRules) SetAttribute(v string) *UpdateLifecyclePolicyRequestDeleteRules {
	s.Attribute = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestDeleteRules) SetThreshold(v string) *UpdateLifecyclePolicyRequestDeleteRules {
	s.Threshold = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestDeleteRules) Validate() error {
	return dara.Validate(s)
}

type UpdateLifecyclePolicyRequestRetrieveRules struct {
	// The attribute of the rule. Valid values:
	//
	// - RetrieveType: the retrieval method.
	//
	// example:
	//
	// RetrieveType
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The threshold of the rule. Valid values for RetrieveType:
	//
	// - AfterVisit: supported when LifecyclePolicyType is set to Auto. This value indicates best-effort recall on visit.
	//
	// - All: supported when LifecyclePolicyType is set to OnDemand. This value indicates that all data is retrieved.
	//
	// example:
	//
	// All
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLifecyclePolicyRequestRetrieveRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyRequestRetrieveRules) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) GetAttribute() *string {
	return s.Attribute
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) GetThreshold() *string {
	return s.Threshold
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) SetAttribute(v string) *UpdateLifecyclePolicyRequestRetrieveRules {
	s.Attribute = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) SetThreshold(v string) *UpdateLifecyclePolicyRequestRetrieveRules {
	s.Threshold = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestRetrieveRules) Validate() error {
	return dara.Validate(s)
}

type UpdateLifecyclePolicyRequestTransitRules struct {
	// The attribute of the rule.
	//
	// Valid values:
	//
	// - Atime: the last access time of the file.
	//
	// example:
	//
	// Atime
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The threshold of the rule.
	//
	// Valid values:
	//
	// - If Attribute is set to Atime, the value indicates the number of days that a file has not been accessed. Valid values: 1 to 365.
	//
	// example:
	//
	// 3
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateLifecyclePolicyRequestTransitRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyRequestTransitRules) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyRequestTransitRules) GetAttribute() *string {
	return s.Attribute
}

func (s *UpdateLifecyclePolicyRequestTransitRules) GetThreshold() *string {
	return s.Threshold
}

func (s *UpdateLifecyclePolicyRequestTransitRules) SetAttribute(v string) *UpdateLifecyclePolicyRequestTransitRules {
	s.Attribute = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestTransitRules) SetThreshold(v string) *UpdateLifecyclePolicyRequestTransitRules {
	s.Threshold = &v
	return s
}

func (s *UpdateLifecyclePolicyRequestTransitRules) Validate() error {
	return dara.Validate(s)
}
