// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceRequirements interface {
	dara.Model
	String() string
	GoString() string
	SetLimits(v map[string]*string) *ResourceRequirements
	GetLimits() map[string]*string
	SetRequests(v map[string]*string) *ResourceRequirements
	GetRequests() map[string]*string
}

type ResourceRequirements struct {
	// The resource limit.
	Limits map[string]*string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	// The resource request.
	Requests map[string]*string `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s ResourceRequirements) String() string {
	return dara.Prettify(s)
}

func (s ResourceRequirements) GoString() string {
	return s.String()
}

func (s *ResourceRequirements) GetLimits() map[string]*string {
	return s.Limits
}

func (s *ResourceRequirements) GetRequests() map[string]*string {
	return s.Requests
}

func (s *ResourceRequirements) SetLimits(v map[string]*string) *ResourceRequirements {
	s.Limits = v
	return s
}

func (s *ResourceRequirements) SetRequests(v map[string]*string) *ResourceRequirements {
	s.Requests = v
	return s
}

func (s *ResourceRequirements) Validate() error {
	return dara.Validate(s)
}
