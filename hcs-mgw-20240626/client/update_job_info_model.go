// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobInfo interface {
	dara.Model
	String() string
	GoString() string
	SetImportQos(v *ImportQos) *UpdateJobInfo
	GetImportQos() *ImportQos
	SetStatus(v string) *UpdateJobInfo
	GetStatus() *string
}

type UpdateJobInfo struct {
	ImportQos *ImportQos `json:"ImportQos,omitempty" xml:"ImportQos,omitempty"`
	// example:
	//
	// IMPORT_JOB_LAUNCHING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateJobInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobInfo) GoString() string {
	return s.String()
}

func (s *UpdateJobInfo) GetImportQos() *ImportQos {
	return s.ImportQos
}

func (s *UpdateJobInfo) GetStatus() *string {
	return s.Status
}

func (s *UpdateJobInfo) SetImportQos(v *ImportQos) *UpdateJobInfo {
	s.ImportQos = v
	return s
}

func (s *UpdateJobInfo) SetStatus(v string) *UpdateJobInfo {
	s.Status = &v
	return s
}

func (s *UpdateJobInfo) Validate() error {
	if s.ImportQos != nil {
		if err := s.ImportQos.Validate(); err != nil {
			return err
		}
	}
	return nil
}
