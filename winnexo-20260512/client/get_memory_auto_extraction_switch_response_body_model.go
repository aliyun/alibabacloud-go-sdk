// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryAutoExtractionSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryAutoExtractionSwitchResponseBody
	GetCode() *string
	SetEnabled(v bool) *GetMemoryAutoExtractionSwitchResponseBody
	GetEnabled() *bool
	SetMessage(v string) *GetMemoryAutoExtractionSwitchResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMemoryAutoExtractionSwitchResponseBody
	GetRequestId() *string
}

type GetMemoryAutoExtractionSwitchResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status of the automatic memory collection switch for the current memory scope. A value of true indicates that the switch is enabled, and a value of false indicates that the switch is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The error description.
	//
	// example:
	//
	// The current user is not an active member of the group cg_12345
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMemoryAutoExtractionSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryAutoExtractionSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetCode(v string) *GetMemoryAutoExtractionSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetEnabled(v bool) *GetMemoryAutoExtractionSwitchResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetMessage(v string) *GetMemoryAutoExtractionSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) SetRequestId(v string) *GetMemoryAutoExtractionSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryAutoExtractionSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
