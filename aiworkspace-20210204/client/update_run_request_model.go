// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*Label) *UpdateRunRequest
	GetLabels() []*Label
	SetName(v string) *UpdateRunRequest
	GetName() *string
	SetParams(v []*RunParam) *UpdateRunRequest
	GetParams() []*RunParam
}

type UpdateRunRequest struct {
	// The labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The run name. The name must meet the following requirements:
	//
	// 	- The name must start with a letter.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be 1 to 63 characters in length.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters.
	Params []*RunParam `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s UpdateRunRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRunRequest) GoString() string {
	return s.String()
}

func (s *UpdateRunRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *UpdateRunRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRunRequest) GetParams() []*RunParam {
	return s.Params
}

func (s *UpdateRunRequest) SetLabels(v []*Label) *UpdateRunRequest {
	s.Labels = v
	return s
}

func (s *UpdateRunRequest) SetName(v string) *UpdateRunRequest {
	s.Name = &v
	return s
}

func (s *UpdateRunRequest) SetParams(v []*RunParam) *UpdateRunRequest {
	s.Params = v
	return s
}

func (s *UpdateRunRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
