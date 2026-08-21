// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *DeleteAlertDestinationRequest
	GetXDebugId() *string
	SetId(v int32) *DeleteAlertDestinationRequest
	GetId() *int32
	SetXSysomInvokeSource(v string) *DeleteAlertDestinationRequest
	GetXSysomInvokeSource() *string
}

type DeleteAlertDestinationRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The ID of the alert contact.
	//
	// example:
	//
	// 1
	Id                 *int32  `json:"id,omitempty" xml:"id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s DeleteAlertDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertDestinationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertDestinationRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *DeleteAlertDestinationRequest) GetId() *int32 {
	return s.Id
}

func (s *DeleteAlertDestinationRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *DeleteAlertDestinationRequest) SetXDebugId(v string) *DeleteAlertDestinationRequest {
	s.XDebugId = &v
	return s
}

func (s *DeleteAlertDestinationRequest) SetId(v int32) *DeleteAlertDestinationRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlertDestinationRequest) SetXSysomInvokeSource(v string) *DeleteAlertDestinationRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *DeleteAlertDestinationRequest) Validate() error {
	return dara.Validate(s)
}
