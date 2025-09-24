// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddImageLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*AddImageLabelsRequestLabels) *AddImageLabelsRequest
	GetLabels() []*AddImageLabelsRequestLabels
}

type AddImageLabelsRequest struct {
	// The list of image tags.
	//
	// This parameter is required.
	Labels []*AddImageLabelsRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s AddImageLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *AddImageLabelsRequest) GetLabels() []*AddImageLabelsRequestLabels {
	return s.Labels
}

func (s *AddImageLabelsRequest) SetLabels(v []*AddImageLabelsRequestLabels) *AddImageLabelsRequest {
	s.Labels = v
	return s
}

func (s *AddImageLabelsRequest) Validate() error {
	return dara.Validate(s)
}

type AddImageLabelsRequestLabels struct {
	// The tag key. The following keys can be added:
	//
	// 	- system.chipType
	//
	// 	- system.dsw.cudaVersion
	//
	// 	- system.dsw.fromImageId
	//
	// 	- system.dsw.fromInstanceId
	//
	// 	- system.dsw.id
	//
	// 	- system.dsw.os
	//
	// 	- system.dsw.osVersion
	//
	// 	- system.dsw.resourceType
	//
	// 	- system.dsw.rootImageId
	//
	// 	- system.dsw.stage
	//
	// 	- system.dsw.tag
	//
	// 	- system.dsw.type
	//
	// 	- system.framework
	//
	// 	- system.origin
	//
	// 	- system.pythonVersion
	//
	// 	- system.source
	//
	// 	- system.supported.dlc
	//
	// 	- system.supported.dsw
	//
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// GPU
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddImageLabelsRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s AddImageLabelsRequestLabels) GoString() string {
	return s.String()
}

func (s *AddImageLabelsRequestLabels) GetKey() *string {
	return s.Key
}

func (s *AddImageLabelsRequestLabels) GetValue() *string {
	return s.Value
}

func (s *AddImageLabelsRequestLabels) SetKey(v string) *AddImageLabelsRequestLabels {
	s.Key = &v
	return s
}

func (s *AddImageLabelsRequestLabels) SetValue(v string) *AddImageLabelsRequestLabels {
	s.Value = &v
	return s
}

func (s *AddImageLabelsRequestLabels) Validate() error {
	return dara.Validate(s)
}
