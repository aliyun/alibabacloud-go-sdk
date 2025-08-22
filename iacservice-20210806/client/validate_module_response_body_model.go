// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ValidateModuleResponseBody
	GetMessage() *string
	SetModuleValidationId(v string) *ValidateModuleResponseBody
	GetModuleValidationId() *string
	SetRequestId(v string) *ValidateModuleResponseBody
	GetRequestId() *string
	SetStatus(v string) *ValidateModuleResponseBody
	GetStatus() *string
}

type ValidateModuleResponseBody struct {
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// BF72A6FB-B07
	ModuleValidationId *string `json:"moduleValidationId,omitempty" xml:"moduleValidationId,omitempty"`
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Validating
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ValidateModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateModuleResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateModuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValidateModuleResponseBody) GetModuleValidationId() *string {
	return s.ModuleValidationId
}

func (s *ValidateModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateModuleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ValidateModuleResponseBody) SetMessage(v string) *ValidateModuleResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateModuleResponseBody) SetModuleValidationId(v string) *ValidateModuleResponseBody {
	s.ModuleValidationId = &v
	return s
}

func (s *ValidateModuleResponseBody) SetRequestId(v string) *ValidateModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateModuleResponseBody) SetStatus(v string) *ValidateModuleResponseBody {
	s.Status = &v
	return s
}

func (s *ValidateModuleResponseBody) Validate() error {
	return dara.Validate(s)
}
