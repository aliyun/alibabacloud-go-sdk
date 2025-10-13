// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyStorageRequirementsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListIntegrationPolicyStorageRequirementsResponseBody
	GetRequestId() *string
	SetStorageRequirements(v []*ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) *ListIntegrationPolicyStorageRequirementsResponseBody
	GetStorageRequirements() []*ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements
}

type ListIntegrationPolicyStorageRequirementsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId           *string                                                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
	StorageRequirements []*ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements `json:"storageRequirements,omitempty" xml:"storageRequirements,omitempty" type:"Repeated"`
}

func (s ListIntegrationPolicyStorageRequirementsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyStorageRequirementsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBody) GetStorageRequirements() []*ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements {
	return s.StorageRequirements
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBody) SetRequestId(v string) *ListIntegrationPolicyStorageRequirementsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBody) SetStorageRequirements(v []*ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) *ListIntegrationPolicyStorageRequirementsResponseBody {
	s.StorageRequirements = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBody) Validate() error {
	if s.StorageRequirements != nil {
		for _, item := range s.StorageRequirements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements struct {
	AddonReleaseNames []*string `json:"addonReleaseNames,omitempty" xml:"addonReleaseNames,omitempty" type:"Repeated"`
	// example:
	//
	// v1
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// example:
	//
	// Pod
	Kind     *string                                                                          `json:"kind,omitempty" xml:"kind,omitempty"`
	Metadata *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	Spec     *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec     `json:"spec,omitempty" xml:"spec,omitempty" type:"Struct"`
	Status   *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus   `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) GetAddonReleaseNames() []*string {
	return s.AddonReleaseNames
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) GetKind() *string {
	return s.Kind
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) GetMetadata() *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata {
	return s.Metadata
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) GetSpec() *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	return s.Spec
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) GetStatus() *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	return s.Status
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) SetAddonReleaseNames(v []*string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements {
	s.AddonReleaseNames = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) SetApiVersion(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements {
	s.ApiVersion = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) SetKind(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements {
	s.Kind = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) SetMetadata(v *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements {
	s.Metadata = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) SetSpec(v *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements {
	s.Spec = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) SetStatus(v *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements {
	s.Status = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirements) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata struct {
	Annotations map[string]*string `json:"annotations,omitempty" xml:"annotations,omitempty"`
	Labels      map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// pod-1234567
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// arms-prom
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) GetAnnotations() map[string]*string {
	return s.Annotations
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) SetAnnotations(v map[string]*string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata {
	s.Annotations = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) SetLabels(v map[string]*string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata {
	s.Labels = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) SetName(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) SetNamespace(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata {
	s.Namespace = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsMetadata) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec struct {
	// example:
	//
	// es-xxxxx
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// example:
	//
	// i-bp122p85gthbniw8rsu9
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// category_predict
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// example:
	//
	// datagrid_cdm
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Region
	ShareScope *string `json:"shareScope,omitempty" xml:"shareScope,omitempty"`
	// example:
	//
	// Prometheus
	StorageType *string            `json:"storageType,omitempty" xml:"storageType,omitempty"`
	SystemTags  map[string]*string `json:"systemTags,omitempty" xml:"systemTags,omitempty"`
	Tags        map[string]*string `json:"tags,omitempty" xml:"tags,omitempty"`
	// example:
	//
	// 12345678
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// test-api
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetEntityId() *string {
	return s.EntityId
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetInstance() *string {
	return s.Instance
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetProject() *string {
	return s.Project
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetRegion() *string {
	return s.Region
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetShareScope() *string {
	return s.ShareScope
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetStorageType() *string {
	return s.StorageType
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetSystemTags() map[string]*string {
	return s.SystemTags
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetTags() map[string]*string {
	return s.Tags
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetUserId() *string {
	return s.UserId
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetEntityId(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.EntityId = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetInstance(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.Instance = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetInstanceName(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.InstanceName = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetProject(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.Project = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetRegion(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.Region = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetShareScope(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.ShareScope = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetStorageType(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.StorageType = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetSystemTags(v map[string]*string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.SystemTags = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetTags(v map[string]*string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.Tags = v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetUserId(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.UserId = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) SetWorkspace(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec {
	s.Workspace = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsSpec) Validate() error {
	return dara.Validate(s)
}

type ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus struct {
	// example:
	//
	// rmq-cn-uqm3ket1t0u
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// http://192.168.xxxxxx
	InterUrl *string `json:"interUrl,omitempty" xml:"interUrl,omitempty"`
	// example:
	//
	// http://100.100.xxxxxx
	IntraUrl *string `json:"intraUrl,omitempty" xml:"intraUrl,omitempty"`
	// example:
	//
	// sr-xxxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// jiuwu_algo
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// xxxx
	PromMetricStore *string `json:"promMetricStore,omitempty" xml:"promMetricStore,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Prometheus
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// example:
	//
	// default
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetInterUrl() *string {
	return s.InterUrl
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetIntraUrl() *string {
	return s.IntraUrl
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetName() *string {
	return s.Name
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetProject() *string {
	return s.Project
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetPromMetricStore() *string {
	return s.PromMetricStore
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetRegion() *string {
	return s.Region
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetStorageType() *string {
	return s.StorageType
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetInstanceId(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.InstanceId = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetInterUrl(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.InterUrl = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetIntraUrl(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.IntraUrl = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetName(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.Name = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetProject(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.Project = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetPromMetricStore(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.PromMetricStore = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetRegion(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.Region = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetStorageType(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.StorageType = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) SetWorkspace(v string) *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus {
	s.Workspace = &v
	return s
}

func (s *ListIntegrationPolicyStorageRequirementsResponseBodyStorageRequirementsStatus) Validate() error {
	return dara.Validate(s)
}
