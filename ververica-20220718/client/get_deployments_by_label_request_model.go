// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreJobSummary(v bool) *GetDeploymentsByLabelRequest
	GetIgnoreJobSummary() *bool
	SetIgnoreResourceSetting(v bool) *GetDeploymentsByLabelRequest
	GetIgnoreResourceSetting() *bool
	SetLabelKey(v string) *GetDeploymentsByLabelRequest
	GetLabelKey() *string
	SetLabelValue(v string) *GetDeploymentsByLabelRequest
	GetLabelValue() *string
}

type GetDeploymentsByLabelRequest struct {
	// Specifies whether to exclude job summary information, such as jobName and status, from the response. If set to true, the response includes only the JobId. This improves performance.
	//
	// example:
	//
	// true
	IgnoreJobSummary *bool `json:"ignoreJobSummary,omitempty" xml:"ignoreJobSummary,omitempty"`
	// Specifies whether to exclude resource configuration information, such as parallelism and the number of CUs, from the response. This reduces the size of the returned data.
	//
	// example:
	//
	// true
	IgnoreResourceSetting *bool `json:"ignoreResourceSetting,omitempty" xml:"ignoreResourceSetting,omitempty"`
	// The label key used for filtering.
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	LabelKey *string `json:"labelKey,omitempty" xml:"labelKey,omitempty"`
	// The label value. You can specify multiple values separated by commas (,) to create an OR condition.
	//
	// This parameter is required.
	//
	// example:
	//
	// value
	LabelValue *string `json:"labelValue,omitempty" xml:"labelValue,omitempty"`
}

func (s GetDeploymentsByLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByLabelRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByLabelRequest) GetIgnoreJobSummary() *bool {
	return s.IgnoreJobSummary
}

func (s *GetDeploymentsByLabelRequest) GetIgnoreResourceSetting() *bool {
	return s.IgnoreResourceSetting
}

func (s *GetDeploymentsByLabelRequest) GetLabelKey() *string {
	return s.LabelKey
}

func (s *GetDeploymentsByLabelRequest) GetLabelValue() *string {
	return s.LabelValue
}

func (s *GetDeploymentsByLabelRequest) SetIgnoreJobSummary(v bool) *GetDeploymentsByLabelRequest {
	s.IgnoreJobSummary = &v
	return s
}

func (s *GetDeploymentsByLabelRequest) SetIgnoreResourceSetting(v bool) *GetDeploymentsByLabelRequest {
	s.IgnoreResourceSetting = &v
	return s
}

func (s *GetDeploymentsByLabelRequest) SetLabelKey(v string) *GetDeploymentsByLabelRequest {
	s.LabelKey = &v
	return s
}

func (s *GetDeploymentsByLabelRequest) SetLabelValue(v string) *GetDeploymentsByLabelRequest {
	s.LabelValue = &v
	return s
}

func (s *GetDeploymentsByLabelRequest) Validate() error {
	return dara.Validate(s)
}
