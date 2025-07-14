// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVersionStatus interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *VersionStatus
	GetErrorMessage() *string
	SetStatus(v string) *VersionStatus
	GetStatus() *string
}

type VersionStatus struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s VersionStatus) String() string {
	return dara.Prettify(s)
}

func (s VersionStatus) GoString() string {
	return s.String()
}

func (s *VersionStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *VersionStatus) GetStatus() *string {
	return s.Status
}

func (s *VersionStatus) SetErrorMessage(v string) *VersionStatus {
	s.ErrorMessage = &v
	return s
}

func (s *VersionStatus) SetStatus(v string) *VersionStatus {
	s.Status = &v
	return s
}

func (s *VersionStatus) Validate() error {
	return dara.Validate(s)
}
