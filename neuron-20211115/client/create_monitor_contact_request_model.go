// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorContact) *CreateMonitorContactRequest
	GetBody() *MonitorContact
}

type CreateMonitorContactRequest struct {
	// This parameter is required.
	Body *MonitorContact `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorContactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorContactRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorContactRequest) GetBody() *MonitorContact {
	return s.Body
}

func (s *CreateMonitorContactRequest) SetBody(v *MonitorContact) *CreateMonitorContactRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorContactRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
