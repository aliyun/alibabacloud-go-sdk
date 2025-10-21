// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfArtifactResult interface {
	dara.Model
	String() string
	GoString() string
	SetCollidingClasses(v []*UdfClass) *UpdateUdfArtifactResult
	GetCollidingClasses() []*UdfClass
	SetMessage(v string) *UpdateUdfArtifactResult
	GetMessage() *string
	SetMissingClasses(v []*UdfClass) *UpdateUdfArtifactResult
	GetMissingClasses() []*UdfClass
	SetUdfArtifact(v *UdfArtifact) *UpdateUdfArtifactResult
	GetUdfArtifact() *UdfArtifact
	SetUpdateSuccess(v bool) *UpdateUdfArtifactResult
	GetUpdateSuccess() *bool
}

type UpdateUdfArtifactResult struct {
	CollidingClasses []*UdfClass  `json:"collidingClasses,omitempty" xml:"collidingClasses,omitempty" type:"Repeated"`
	Message          *string      `json:"message,omitempty" xml:"message,omitempty"`
	MissingClasses   []*UdfClass  `json:"missingClasses,omitempty" xml:"missingClasses,omitempty" type:"Repeated"`
	UdfArtifact      *UdfArtifact `json:"udfArtifact,omitempty" xml:"udfArtifact,omitempty"`
	UpdateSuccess    *bool        `json:"updateSuccess,omitempty" xml:"updateSuccess,omitempty"`
}

func (s UpdateUdfArtifactResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfArtifactResult) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactResult) GetCollidingClasses() []*UdfClass {
	return s.CollidingClasses
}

func (s *UpdateUdfArtifactResult) GetMessage() *string {
	return s.Message
}

func (s *UpdateUdfArtifactResult) GetMissingClasses() []*UdfClass {
	return s.MissingClasses
}

func (s *UpdateUdfArtifactResult) GetUdfArtifact() *UdfArtifact {
	return s.UdfArtifact
}

func (s *UpdateUdfArtifactResult) GetUpdateSuccess() *bool {
	return s.UpdateSuccess
}

func (s *UpdateUdfArtifactResult) SetCollidingClasses(v []*UdfClass) *UpdateUdfArtifactResult {
	s.CollidingClasses = v
	return s
}

func (s *UpdateUdfArtifactResult) SetMessage(v string) *UpdateUdfArtifactResult {
	s.Message = &v
	return s
}

func (s *UpdateUdfArtifactResult) SetMissingClasses(v []*UdfClass) *UpdateUdfArtifactResult {
	s.MissingClasses = v
	return s
}

func (s *UpdateUdfArtifactResult) SetUdfArtifact(v *UdfArtifact) *UpdateUdfArtifactResult {
	s.UdfArtifact = v
	return s
}

func (s *UpdateUdfArtifactResult) SetUpdateSuccess(v bool) *UpdateUdfArtifactResult {
	s.UpdateSuccess = &v
	return s
}

func (s *UpdateUdfArtifactResult) Validate() error {
	if s.CollidingClasses != nil {
		for _, item := range s.CollidingClasses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MissingClasses != nil {
		for _, item := range s.MissingClasses {
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
