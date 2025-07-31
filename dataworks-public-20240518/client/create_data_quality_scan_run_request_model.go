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
	DataQualityScanId *int64                                          `json:"DataQualityScanId,omitempty" xml:"DataQualityScanId,omitempty"`
	Parameters        []*CreateDataQualityScanRunRequestParameters    `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	ProjectId         *int64                                          `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RuntimeResource   *CreateDataQualityScanRunRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateDataQualityScanRunRequestParameters struct {
	// example:
	//
	// regiondt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Cu    *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	Id    *string  `json:"Id,omitempty" xml:"Id,omitempty"`
	Image *string  `json:"Image,omitempty" xml:"Image,omitempty"`
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
