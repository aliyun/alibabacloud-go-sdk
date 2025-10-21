// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfArtifactResult interface {
	dara.Model
	String() string
	GoString() string
	SetCollidingClasses(v []*UdfClass) *CreateUdfArtifactResult
	GetCollidingClasses() []*UdfClass
	SetCreateSuccess(v bool) *CreateUdfArtifactResult
	GetCreateSuccess() *bool
	SetMessage(v string) *CreateUdfArtifactResult
	GetMessage() *string
	SetUdfArtifact(v *UdfArtifact) *CreateUdfArtifactResult
	GetUdfArtifact() *UdfArtifact
}

type CreateUdfArtifactResult struct {
	CollidingClasses []*UdfClass  `json:"collidingClasses,omitempty" xml:"collidingClasses,omitempty" type:"Repeated"`
	CreateSuccess    *bool        `json:"createSuccess,omitempty" xml:"createSuccess,omitempty"`
	Message          *string      `json:"message,omitempty" xml:"message,omitempty"`
	UdfArtifact      *UdfArtifact `json:"udfArtifact,omitempty" xml:"udfArtifact,omitempty"`
}

func (s CreateUdfArtifactResult) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfArtifactResult) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactResult) GetCollidingClasses() []*UdfClass {
	return s.CollidingClasses
}

func (s *CreateUdfArtifactResult) GetCreateSuccess() *bool {
	return s.CreateSuccess
}

func (s *CreateUdfArtifactResult) GetMessage() *string {
	return s.Message
}

func (s *CreateUdfArtifactResult) GetUdfArtifact() *UdfArtifact {
	return s.UdfArtifact
}

func (s *CreateUdfArtifactResult) SetCollidingClasses(v []*UdfClass) *CreateUdfArtifactResult {
	s.CollidingClasses = v
	return s
}

func (s *CreateUdfArtifactResult) SetCreateSuccess(v bool) *CreateUdfArtifactResult {
	s.CreateSuccess = &v
	return s
}

func (s *CreateUdfArtifactResult) SetMessage(v string) *CreateUdfArtifactResult {
	s.Message = &v
	return s
}

func (s *CreateUdfArtifactResult) SetUdfArtifact(v *UdfArtifact) *CreateUdfArtifactResult {
	s.UdfArtifact = v
	return s
}

func (s *CreateUdfArtifactResult) Validate() error {
	if s.CollidingClasses != nil {
		for _, item := range s.CollidingClasses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UdfArtifact != nil {
		if err := s.UdfArtifact.Validate(); err != nil {
			return err
		}
	}
	return nil
}
