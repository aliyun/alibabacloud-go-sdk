// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertDestinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetAlertDestinationRequest
	GetXDebugId() *string
	SetId(v int32) *GetAlertDestinationRequest
	GetId() *int32
	SetXSysomInvokeSource(v string) *GetAlertDestinationRequest
	GetXSysomInvokeSource() *string
}

type GetAlertDestinationRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The alert contact ID.
	//
	// example:
	//
	// 1
	Id                 *int32  `json:"id,omitempty" xml:"id,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s GetAlertDestinationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlertDestinationRequest) GoString() string {
	return s.String()
}

func (s *GetAlertDestinationRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *GetAlertDestinationRequest) GetId() *int32 {
	return s.Id
}

func (s *GetAlertDestinationRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetAlertDestinationRequest) SetXDebugId(v string) *GetAlertDestinationRequest {
	s.XDebugId = &v
	return s
}

func (s *GetAlertDestinationRequest) SetId(v int32) *GetAlertDestinationRequest {
	s.Id = &v
	return s
}

func (s *GetAlertDestinationRequest) SetXSysomInvokeSource(v string) *GetAlertDestinationRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetAlertDestinationRequest) Validate() error {
	return dara.Validate(s)
}
