// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebVersionStatus interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *WebVersionStatus
	GetErrorMessage() *string
	SetStatus(v string) *WebVersionStatus
	GetStatus() *string
}

type WebVersionStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s WebVersionStatus) String() string {
	return dara.Prettify(s)
}

func (s WebVersionStatus) GoString() string {
	return s.String()
}

func (s *WebVersionStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *WebVersionStatus) GetStatus() *string {
	return s.Status
}

func (s *WebVersionStatus) SetErrorMessage(v string) *WebVersionStatus {
	s.ErrorMessage = &v
	return s
}

func (s *WebVersionStatus) SetStatus(v string) *WebVersionStatus {
	s.Status = &v
	return s
}

func (s *WebVersionStatus) Validate() error {
	return dara.Validate(s)
}
