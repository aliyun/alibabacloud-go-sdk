// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationWithInstanceCount interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCount(v int64) *WebApplicationWithInstanceCount
	GetInstanceCount() *int64
	SetWebApplication(v *WebApplication) *WebApplicationWithInstanceCount
	GetWebApplication() *WebApplication
}

type WebApplicationWithInstanceCount struct {
	InstanceCount  *int64          `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	WebApplication *WebApplication `json:"WebApplication,omitempty" xml:"WebApplication,omitempty"`
}

func (s WebApplicationWithInstanceCount) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationWithInstanceCount) GoString() string {
	return s.String()
}

func (s *WebApplicationWithInstanceCount) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *WebApplicationWithInstanceCount) GetWebApplication() *WebApplication {
	return s.WebApplication
}

func (s *WebApplicationWithInstanceCount) SetInstanceCount(v int64) *WebApplicationWithInstanceCount {
	s.InstanceCount = &v
	return s
}

func (s *WebApplicationWithInstanceCount) SetWebApplication(v *WebApplication) *WebApplicationWithInstanceCount {
	s.WebApplication = v
	return s
}

func (s *WebApplicationWithInstanceCount) Validate() error {
	return dara.Validate(s)
}
