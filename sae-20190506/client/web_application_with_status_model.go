// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationWithStatus interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v *WebApplicationStatus) *WebApplicationWithStatus
	GetStatus() *WebApplicationStatus
	SetWebApplication(v *WebApplication) *WebApplicationWithStatus
	GetWebApplication() *WebApplication
}

type WebApplicationWithStatus struct {
	Status         *WebApplicationStatus `json:"Status,omitempty" xml:"Status,omitempty"`
	WebApplication *WebApplication       `json:"WebApplication,omitempty" xml:"WebApplication,omitempty"`
}

func (s WebApplicationWithStatus) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationWithStatus) GoString() string {
	return s.String()
}

func (s *WebApplicationWithStatus) GetStatus() *WebApplicationStatus {
	return s.Status
}

func (s *WebApplicationWithStatus) GetWebApplication() *WebApplication {
	return s.WebApplication
}

func (s *WebApplicationWithStatus) SetStatus(v *WebApplicationStatus) *WebApplicationWithStatus {
	s.Status = v
	return s
}

func (s *WebApplicationWithStatus) SetWebApplication(v *WebApplication) *WebApplicationWithStatus {
	s.WebApplication = v
	return s
}

func (s *WebApplicationWithStatus) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	if s.WebApplication != nil {
		if err := s.WebApplication.Validate(); err != nil {
			return err
		}
	}
	return nil
}
