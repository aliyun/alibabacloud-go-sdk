// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneDefensePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *CreateSceneDefensePolicyRequest
	GetEndTime() *int64
	SetName(v string) *CreateSceneDefensePolicyRequest
	GetName() *string
	SetStartTime(v int64) *CreateSceneDefensePolicyRequest
	GetStartTime() *int64
	SetTemplate(v string) *CreateSceneDefensePolicyRequest
	GetTemplate() *string
}

type CreateSceneDefensePolicyRequest struct {
	// The end time of the policy. This value is a UNIX timestamp. Units: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1586016000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpolicy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The start time of the policy. This value is a UNIX timestamp. Units: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1585670400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The template of the policy. Valid values:
	//
	// 	- **promotion**: important activity.
	//
	// 	- **bypass**: all traffic forwarded.
	//
	// This parameter is required.
	//
	// example:
	//
	// promotion
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s CreateSceneDefensePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSceneDefensePolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateSceneDefensePolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateSceneDefensePolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateSceneDefensePolicyRequest) GetTemplate() *string {
	return s.Template
}

func (s *CreateSceneDefensePolicyRequest) SetEndTime(v int64) *CreateSceneDefensePolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) SetName(v string) *CreateSceneDefensePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) SetStartTime(v int64) *CreateSceneDefensePolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) SetTemplate(v string) *CreateSceneDefensePolicyRequest {
	s.Template = &v
	return s
}

func (s *CreateSceneDefensePolicyRequest) Validate() error {
	return dara.Validate(s)
}
