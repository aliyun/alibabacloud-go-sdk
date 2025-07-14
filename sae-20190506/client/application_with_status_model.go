// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplicationWithStatus interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *Application) *ApplicationWithStatus
	GetApplication() *Application
	SetStatus(v *ApplicationStatus) *ApplicationWithStatus
	GetStatus() *ApplicationStatus
}

type ApplicationWithStatus struct {
	Application *Application       `json:"application,omitempty" xml:"application,omitempty"`
	Status      *ApplicationStatus `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ApplicationWithStatus) String() string {
	return dara.Prettify(s)
}

func (s ApplicationWithStatus) GoString() string {
	return s.String()
}

func (s *ApplicationWithStatus) GetApplication() *Application {
	return s.Application
}

func (s *ApplicationWithStatus) GetStatus() *ApplicationStatus {
	return s.Status
}

func (s *ApplicationWithStatus) SetApplication(v *Application) *ApplicationWithStatus {
	s.Application = v
	return s
}

func (s *ApplicationWithStatus) SetStatus(v *ApplicationStatus) *ApplicationWithStatus {
	s.Status = v
	return s
}

func (s *ApplicationWithStatus) Validate() error {
	return dara.Validate(s)
}
