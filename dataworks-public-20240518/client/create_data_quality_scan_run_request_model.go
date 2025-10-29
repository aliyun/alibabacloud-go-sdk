// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityScanId(v int64) *CreateDataQualityScanRunRequest
	GetDataQualityScanId() *int64
	SetParameters(v []*CreateDataQualityScanRunRequestParameters) *CreateDataQualityScanRunRequest
	GetParameters() []*CreateDataQualityScanRunRequestParameters
	SetProjectId(v int64) *CreateDataQualityScanRunRequest
	GetProjectId() *int64
	SetRuntimeResource(v *CreateDataQualityScanRunRequestRuntimeResource) *CreateDataQualityScanRunRequest
	GetRuntimeResource() *CreateDataQualityScanRunRequestRuntimeResource
}

type CreateDataQualityScanRunRequest struct {
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 20000001
	DataQualityScanId *int64 `json:"DataQualityScanId,omitempty" xml:"DataQualityScanId,omitempty"`
	// The parameter settings used during the actual run. The `triggerTime` parameter is required.
	Parameters []*CreateDataQualityScanRunRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The project ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling resource group used when running the data quality monitor. This resource group uses the same data structure as in the scheduling API.
	RuntimeResource *CreateDataQualityScanRunRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
}

func (s CreateDataQualityScanRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRunRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRunRequest) GetDataQualityScanId() *int64 {
	return s.DataQualityScanId
}

func (s *CreateDataQualityScanRunRequest) GetParameters() []*CreateDataQualityScanRunRequestParameters {
	return s.Parameters
}

func (s *CreateDataQualityScanRunRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityScanRunRequest) GetRuntimeResource() *CreateDataQualityScanRunRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *CreateDataQualityScanRunRequest) SetDataQualityScanId(v int64) *CreateDataQualityScanRunRequest {
	s.DataQualityScanId = &v
	return s
}

func (s *CreateDataQualityScanRunRequest) SetParameters(v []*CreateDataQualityScanRunRequestParameters) *CreateDataQualityScanRunRequest {
	s.Parameters = v
	return s
}

func (s *CreateDataQualityScanRunRequest) SetProjectId(v int64) *CreateDataQualityScanRunRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityScanRunRequest) SetRuntimeResource(v *CreateDataQualityScanRunRequestRuntimeResource) *CreateDataQualityScanRunRequest {
	s.RuntimeResource = v
	return s
}

func (s *CreateDataQualityScanRunRequest) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataQualityScanRunRequestParameters struct {
	// The parameter name.
	//
	// example:
	//
	// regiondt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value. You can use a scheduling time expression.
	//
	// example:
	//
	// cn-shanghai$[yyyy-mm-dd-1]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataQualityScanRunRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRunRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRunRequestParameters) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityScanRunRequestParameters) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityScanRunRequestParameters) SetName(v string) *CreateDataQualityScanRunRequestParameters {
	s.Name = &v
	return s
}

func (s *CreateDataQualityScanRunRequestParameters) SetValue(v string) *CreateDataQualityScanRunRequestParameters {
	s.Value = &v
	return s
}

func (s *CreateDataQualityScanRunRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityScanRunRequestRuntimeResource struct {
	// The Compute Resources (CUs) reserved for running the data quality monitor in the resource group.
	//
	// example:
	//
	// 0.25
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// e9455a13-ff00-4965-833c-337546ba4854
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image settings used when running the data quality monitor in the resource group.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s CreateDataQualityScanRunRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRunRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRunRequestRuntimeResource) GetCu() *float32 {
	return s.Cu
}

func (s *CreateDataQualityScanRunRequestRuntimeResource) GetId() *string {
	return s.Id
}

func (s *CreateDataQualityScanRunRequestRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *CreateDataQualityScanRunRequestRuntimeResource) SetCu(v float32) *CreateDataQualityScanRunRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *CreateDataQualityScanRunRequestRuntimeResource) SetId(v string) *CreateDataQualityScanRunRequestRuntimeResource {
	s.Id = &v
	return s
}

func (s *CreateDataQualityScanRunRequestRuntimeResource) SetImage(v string) *CreateDataQualityScanRunRequestRuntimeResource {
	s.Image = &v
	return s
}

func (s *CreateDataQualityScanRunRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}
