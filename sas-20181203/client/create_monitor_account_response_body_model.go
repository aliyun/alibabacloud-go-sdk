// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMonitorAccountResponseBody
	GetRequestId() *string
}

type CreateMonitorAccountResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMonitorAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitorAccountResponseBody) SetRequestId(v string) *CreateMonitorAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
