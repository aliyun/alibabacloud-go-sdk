// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanRunShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityScanId(v int64) *CreateDataQualityScanRunShrinkRequest
	GetDataQualityScanId() *int64
	SetParametersShrink(v string) *CreateDataQualityScanRunShrinkRequest
	GetParametersShrink() *string
	SetProjectId(v int64) *CreateDataQualityScanRunShrinkRequest
	GetProjectId() *int64
	SetRuntimeResourceShrink(v string) *CreateDataQualityScanRunShrinkRequest
	GetRuntimeResourceShrink() *string
}

type CreateDataQualityScanRunShrinkRequest struct {
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 20000001
	DataQualityScanId *int64 `json:"DataQualityScanId,omitempty" xml:"DataQualityScanId,omitempty"`
	// The parameter settings used during the actual run. The `triggerTime` parameter is required.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling resource group used when running the data quality monitor. This resource group uses the same data structure as in the scheduling API.
	RuntimeResourceShrink *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
}

func (s CreateDataQualityScanRunShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRunShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRunShrinkRequest) GetDataQualityScanId() *int64 {
	return s.DataQualityScanId
}

func (s *CreateDataQualityScanRunShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateDataQualityScanRunShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityScanRunShrinkRequest) GetRuntimeResourceShrink() *string {
	return s.RuntimeResourceShrink
}

func (s *CreateDataQualityScanRunShrinkRequest) SetDataQualityScanId(v int64) *CreateDataQualityScanRunShrinkRequest {
	s.DataQualityScanId = &v
	return s
}

func (s *CreateDataQualityScanRunShrinkRequest) SetParametersShrink(v string) *CreateDataQualityScanRunShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateDataQualityScanRunShrinkRequest) SetProjectId(v int64) *CreateDataQualityScanRunShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityScanRunShrinkRequest) SetRuntimeResourceShrink(v string) *CreateDataQualityScanRunShrinkRequest {
	s.RuntimeResourceShrink = &v
	return s
}

func (s *CreateDataQualityScanRunShrinkRequest) Validate() error {
	return dara.Validate(s)
}
