// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfArtifactResult interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteSuccess(v bool) *DeleteUdfArtifactResult
	GetDeleteSuccess() *bool
	SetMessage(v string) *DeleteUdfArtifactResult
	GetMessage() *string
	SetReferencedClasses(v []*UdfClass) *DeleteUdfArtifactResult
	GetReferencedClasses() []*UdfClass
}

type DeleteUdfArtifactResult struct {
	DeleteSuccess     *bool       `json:"deleteSuccess,omitempty" xml:"deleteSuccess,omitempty"`
	Message           *string     `json:"message,omitempty" xml:"message,omitempty"`
	ReferencedClasses []*UdfClass `json:"referencedClasses,omitempty" xml:"referencedClasses,omitempty" type:"Repeated"`
}

func (s DeleteUdfArtifactResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfArtifactResult) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactResult) GetDeleteSuccess() *bool {
	return s.DeleteSuccess
}

func (s *DeleteUdfArtifactResult) GetMessage() *string {
	return s.Message
}

func (s *DeleteUdfArtifactResult) GetReferencedClasses() []*UdfClass {
	return s.ReferencedClasses
}

func (s *DeleteUdfArtifactResult) SetDeleteSuccess(v bool) *DeleteUdfArtifactResult {
	s.DeleteSuccess = &v
	return s
}

func (s *DeleteUdfArtifactResult) SetMessage(v string) *DeleteUdfArtifactResult {
	s.Message = &v
	return s
}

func (s *DeleteUdfArtifactResult) SetReferencedClasses(v []*UdfClass) *DeleteUdfArtifactResult {
	s.ReferencedClasses = v
	return s
}

func (s *DeleteUdfArtifactResult) Validate() error {
	if s.ReferencedClasses != nil {
		for _, item := range s.ReferencedClasses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
