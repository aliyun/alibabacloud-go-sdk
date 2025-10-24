// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceName(v string) *UpdateResourceRequest
	GetResourceName() *string
	SetSelfManagedResourceOptions(v *UpdateResourceRequestSelfManagedResourceOptions) *UpdateResourceRequest
	GetSelfManagedResourceOptions() *UpdateResourceRequestSelfManagedResourceOptions
}

type UpdateResourceRequest struct {
	// The new name of the resource group after the update. The name can be up to 27 characters in length.
	//
	// example:
	//
	// iot
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The configuration items of the self-managed resource group.
	SelfManagedResourceOptions *UpdateResourceRequestSelfManagedResourceOptions `json:"SelfManagedResourceOptions,omitempty" xml:"SelfManagedResourceOptions,omitempty" type:"Struct"`
}

func (s UpdateResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *UpdateResourceRequest) GetSelfManagedResourceOptions() *UpdateResourceRequestSelfManagedResourceOptions {
	return s.SelfManagedResourceOptions
}

func (s *UpdateResourceRequest) SetResourceName(v string) *UpdateResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *UpdateResourceRequest) SetSelfManagedResourceOptions(v *UpdateResourceRequestSelfManagedResourceOptions) *UpdateResourceRequest {
	s.SelfManagedResourceOptions = v
	return s
}

func (s *UpdateResourceRequest) Validate() error {
	if s.SelfManagedResourceOptions != nil {
		if err := s.SelfManagedResourceOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateResourceRequestSelfManagedResourceOptions struct {
	// Tag tag key-value pairs for nodes.
	NodeMatchLabels map[string]*string `json:"NodeMatchLabels,omitempty" xml:"NodeMatchLabels,omitempty"`
	// Tolerations for nodes.
	NodeTolerations []*UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations `json:"NodeTolerations,omitempty" xml:"NodeTolerations,omitempty" type:"Repeated"`
}

func (s UpdateResourceRequestSelfManagedResourceOptions) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequestSelfManagedResourceOptions) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequestSelfManagedResourceOptions) GetNodeMatchLabels() map[string]*string {
	return s.NodeMatchLabels
}

func (s *UpdateResourceRequestSelfManagedResourceOptions) GetNodeTolerations() []*UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	return s.NodeTolerations
}

func (s *UpdateResourceRequestSelfManagedResourceOptions) SetNodeMatchLabels(v map[string]*string) *UpdateResourceRequestSelfManagedResourceOptions {
	s.NodeMatchLabels = v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptions) SetNodeTolerations(v []*UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) *UpdateResourceRequestSelfManagedResourceOptions {
	s.NodeTolerations = v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptions) Validate() error {
	if s.NodeTolerations != nil {
		for _, item := range s.NodeTolerations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations struct {
	// The effect.
	//
	// Valid values:
	//
	// - PreferNoSchedule
	//
	// - NoSchedule
	//
	// - NoExecute
	//
	// example:
	//
	// NoSchedule
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	// The key name.
	//
	// example:
	//
	// key1
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Relationship between key names and key values.
	//
	// Valid values:
	//
	// - Equal
	//
	// - Exists
	//
	// example:
	//
	// Equal
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The key value.
	//
	// example:
	//
	// value1
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetEffect() *string {
	return s.Effect
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetKey() *string {
	return s.Key
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetOperator() *string {
	return s.Operator
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) GetValue() *string {
	return s.Value
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetEffect(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Effect = &v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetKey(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Key = &v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetOperator(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Operator = &v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) SetValue(v string) *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations {
	s.Value = &v
	return s
}

func (s *UpdateResourceRequestSelfManagedResourceOptionsNodeTolerations) Validate() error {
	return dara.Validate(s)
}
