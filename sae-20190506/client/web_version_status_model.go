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
	// The error message of the application execution. If the execution is successful, a null value is returned.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The state of the application execution. Valid values:
	//
	// 	- **Completed**: The execution is successful.
	//
	// 	- **Updating**:The instance is being updated.
	//
	// 	- **Updating**:The execution failed and a non-null error message is returned.
	//
	// example:
	//
	// Completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
