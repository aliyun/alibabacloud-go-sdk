// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ValidateConnectionResponseBody
	GetCode() *string
	SetDetails(v map[string]*string) *ValidateConnectionResponseBody
	GetDetails() map[string]*string
	SetMessage(v string) *ValidateConnectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *ValidateConnectionResponseBody
	GetRequestId() *string
	SetStatus(v string) *ValidateConnectionResponseBody
	GetStatus() *string
}

type ValidateConnectionResponseBody struct {
	// example:
	//
	// 403
	Code    *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	Details map[string]*string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// The connection is reachable.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ValidateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateConnectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ValidateConnectionResponseBody) GetDetails() map[string]*string {
	return s.Details
}

func (s *ValidateConnectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValidateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateConnectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ValidateConnectionResponseBody) SetCode(v string) *ValidateConnectionResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateConnectionResponseBody) SetDetails(v map[string]*string) *ValidateConnectionResponseBody {
	s.Details = v
	return s
}

func (s *ValidateConnectionResponseBody) SetMessage(v string) *ValidateConnectionResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateConnectionResponseBody) SetRequestId(v string) *ValidateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateConnectionResponseBody) SetStatus(v string) *ValidateConnectionResponseBody {
	s.Status = &v
	return s
}

func (s *ValidateConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
