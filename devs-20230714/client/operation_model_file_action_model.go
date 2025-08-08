// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationModelFileAction interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *OperationModelFileAction
	GetAction() *string
	SetDestination(v string) *OperationModelFileAction
	GetDestination() *string
	SetSource(v string) *OperationModelFileAction
	GetSource() *string
	SetTarget(v string) *OperationModelFileAction
	GetTarget() *string
}

type OperationModelFileAction struct {
	// This parameter is required.
	Action      *string `json:"action,omitempty" xml:"action,omitempty"`
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty"`
	Source      *string `json:"source,omitempty" xml:"source,omitempty"`
	Target      *string `json:"target,omitempty" xml:"target,omitempty"`
}

func (s OperationModelFileAction) String() string {
	return dara.Prettify(s)
}

func (s OperationModelFileAction) GoString() string {
	return s.String()
}

func (s *OperationModelFileAction) GetAction() *string {
	return s.Action
}

func (s *OperationModelFileAction) GetDestination() *string {
	return s.Destination
}

func (s *OperationModelFileAction) GetSource() *string {
	return s.Source
}

func (s *OperationModelFileAction) GetTarget() *string {
	return s.Target
}

func (s *OperationModelFileAction) SetAction(v string) *OperationModelFileAction {
	s.Action = &v
	return s
}

func (s *OperationModelFileAction) SetDestination(v string) *OperationModelFileAction {
	s.Destination = &v
	return s
}

func (s *OperationModelFileAction) SetSource(v string) *OperationModelFileAction {
	s.Source = &v
	return s
}

func (s *OperationModelFileAction) SetTarget(v string) *OperationModelFileAction {
	s.Target = &v
	return s
}

func (s *OperationModelFileAction) Validate() error {
	return dara.Validate(s)
}
