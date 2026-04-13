// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerraformStateDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJob(v *GetTerraformStateDetectionResponseBodyJob) *GetTerraformStateDetectionResponseBody
	GetJob() *GetTerraformStateDetectionResponseBodyJob
	SetRequestId(v string) *GetTerraformStateDetectionResponseBody
	GetRequestId() *string
}

type GetTerraformStateDetectionResponseBody struct {
	Job *GetTerraformStateDetectionResponseBodyJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTerraformStateDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionResponseBody) GetJob() *GetTerraformStateDetectionResponseBodyJob {
	return s.Job
}

func (s *GetTerraformStateDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTerraformStateDetectionResponseBody) SetJob(v *GetTerraformStateDetectionResponseBodyJob) *GetTerraformStateDetectionResponseBody {
	s.Job = v
	return s
}

func (s *GetTerraformStateDetectionResponseBody) SetRequestId(v string) *GetTerraformStateDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBody) Validate() error {
	if s.Job != nil {
		if err := s.Job.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTerraformStateDetectionResponseBodyJob struct {
	ChangedResources []*GetTerraformStateDetectionResponseBodyJobChangedResources `json:"changedResources,omitempty" xml:"changedResources,omitempty" type:"Repeated"`
	DriftedResources []*GetTerraformStateDetectionResponseBodyJobDriftedResources `json:"driftedResources,omitempty" xml:"driftedResources,omitempty" type:"Repeated"`
	// example:
	//
	// planned failed
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// stack-as181axxxxxx:development_xxxx
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// Errored
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// Stack
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTerraformStateDetectionResponseBodyJob) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionResponseBodyJob) GetChangedResources() []*GetTerraformStateDetectionResponseBodyJobChangedResources {
	return s.ChangedResources
}

func (s *GetTerraformStateDetectionResponseBodyJob) GetDriftedResources() []*GetTerraformStateDetectionResponseBodyJobDriftedResources {
	return s.DriftedResources
}

func (s *GetTerraformStateDetectionResponseBodyJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTerraformStateDetectionResponseBodyJob) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetTerraformStateDetectionResponseBodyJob) GetStatus() *string {
	return s.Status
}

func (s *GetTerraformStateDetectionResponseBodyJob) GetType() *string {
	return s.Type
}

func (s *GetTerraformStateDetectionResponseBodyJob) SetChangedResources(v []*GetTerraformStateDetectionResponseBodyJobChangedResources) *GetTerraformStateDetectionResponseBodyJob {
	s.ChangedResources = v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJob) SetDriftedResources(v []*GetTerraformStateDetectionResponseBodyJobDriftedResources) *GetTerraformStateDetectionResponseBodyJob {
	s.DriftedResources = v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJob) SetErrorMessage(v string) *GetTerraformStateDetectionResponseBodyJob {
	s.ErrorMessage = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJob) SetIdentifier(v string) *GetTerraformStateDetectionResponseBodyJob {
	s.Identifier = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJob) SetStatus(v string) *GetTerraformStateDetectionResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJob) SetType(v string) *GetTerraformStateDetectionResponseBodyJob {
	s.Type = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJob) Validate() error {
	if s.ChangedResources != nil {
		for _, item := range s.ChangedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DriftedResources != nil {
		for _, item := range s.DriftedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTerraformStateDetectionResponseBodyJobChangedResources struct {
	AttributeChanges []*GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges `json:"attributeChanges,omitempty" xml:"attributeChanges,omitempty" type:"Repeated"`
	// example:
	//
	// create
	ChangedType *string `json:"changedType,omitempty" xml:"changedType,omitempty"`
	// example:
	//
	// false
	HasDrift *bool `json:"hasDrift,omitempty" xml:"hasDrift,omitempty"`
	// example:
	//
	// vpc-axxxxx
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// vpc:alicloud_vpc.default
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
}

func (s GetTerraformStateDetectionResponseBodyJobChangedResources) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionResponseBodyJobChangedResources) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) GetAttributeChanges() []*GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges {
	return s.AttributeChanges
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) GetChangedType() *string {
	return s.ChangedType
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) GetHasDrift() *bool {
	return s.HasDrift
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) GetResourceIdentifier() *string {
	return s.ResourceIdentifier
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) SetAttributeChanges(v []*GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) *GetTerraformStateDetectionResponseBodyJobChangedResources {
	s.AttributeChanges = v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) SetChangedType(v string) *GetTerraformStateDetectionResponseBodyJobChangedResources {
	s.ChangedType = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) SetHasDrift(v bool) *GetTerraformStateDetectionResponseBodyJobChangedResources {
	s.HasDrift = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) SetResourceId(v string) *GetTerraformStateDetectionResponseBodyJobChangedResources {
	s.ResourceId = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) SetResourceIdentifier(v string) *GetTerraformStateDetectionResponseBodyJobChangedResources {
	s.ResourceIdentifier = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResources) Validate() error {
	if s.AttributeChanges != nil {
		for _, item := range s.AttributeChanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges struct {
	// example:
	//
	// vpc_name
	AttributePath *string `json:"attributePath,omitempty" xml:"attributePath,omitempty"`
	// example:
	//
	// test_remote
	RemoteValue *string `json:"remoteValue,omitempty" xml:"remoteValue,omitempty"`
	// example:
	//
	// test_hcl
	TemplateValue *string `json:"templateValue,omitempty" xml:"templateValue,omitempty"`
}

func (s GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) GetAttributePath() *string {
	return s.AttributePath
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) GetRemoteValue() *string {
	return s.RemoteValue
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) GetTemplateValue() *string {
	return s.TemplateValue
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) SetAttributePath(v string) *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges {
	s.AttributePath = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) SetRemoteValue(v string) *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges {
	s.RemoteValue = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) SetTemplateValue(v string) *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges {
	s.TemplateValue = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobChangedResourcesAttributeChanges) Validate() error {
	return dara.Validate(s)
}

type GetTerraformStateDetectionResponseBodyJobDriftedResources struct {
	AttributeDrifts []*GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts `json:"attributeDrifts,omitempty" xml:"attributeDrifts,omitempty" type:"Repeated"`
	// example:
	//
	// update
	DriftedType *string `json:"driftedType,omitempty" xml:"driftedType,omitempty"`
	// example:
	//
	// vpc-bxxxxx
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// vpc:alicloud_vpc.default2
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
}

func (s GetTerraformStateDetectionResponseBodyJobDriftedResources) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionResponseBodyJobDriftedResources) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) GetAttributeDrifts() []*GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts {
	return s.AttributeDrifts
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) GetDriftedType() *string {
	return s.DriftedType
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) GetResourceIdentifier() *string {
	return s.ResourceIdentifier
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) SetAttributeDrifts(v []*GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) *GetTerraformStateDetectionResponseBodyJobDriftedResources {
	s.AttributeDrifts = v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) SetDriftedType(v string) *GetTerraformStateDetectionResponseBodyJobDriftedResources {
	s.DriftedType = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) SetResourceId(v string) *GetTerraformStateDetectionResponseBodyJobDriftedResources {
	s.ResourceId = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) SetResourceIdentifier(v string) *GetTerraformStateDetectionResponseBodyJobDriftedResources {
	s.ResourceIdentifier = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResources) Validate() error {
	if s.AttributeDrifts != nil {
		for _, item := range s.AttributeDrifts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts struct {
	// example:
	//
	// vpc_name
	AttributePath *string `json:"attributePath,omitempty" xml:"attributePath,omitempty"`
	// example:
	//
	// test_remote
	RemoteValue *string `json:"remoteValue,omitempty" xml:"remoteValue,omitempty"`
	// example:
	//
	// test_state
	StateValue *string `json:"stateValue,omitempty" xml:"stateValue,omitempty"`
}

func (s GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) GetAttributePath() *string {
	return s.AttributePath
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) GetRemoteValue() *string {
	return s.RemoteValue
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) GetStateValue() *string {
	return s.StateValue
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) SetAttributePath(v string) *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts {
	s.AttributePath = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) SetRemoteValue(v string) *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts {
	s.RemoteValue = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) SetStateValue(v string) *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts {
	s.StateValue = &v
	return s
}

func (s *GetTerraformStateDetectionResponseBodyJobDriftedResourcesAttributeDrifts) Validate() error {
	return dara.Validate(s)
}
