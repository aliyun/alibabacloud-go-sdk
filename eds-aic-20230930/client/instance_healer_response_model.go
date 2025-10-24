// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceHealerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstanceHealerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstanceHealerResponse
	GetStatusCode() *int32
	SetBody(v *InstanceHealerResponseBody) *InstanceHealerResponse
	GetBody() *InstanceHealerResponseBody
}

type InstanceHealerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstanceHealerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstanceHealerResponse) String() string {
	return dara.Prettify(s)
}

func (s InstanceHealerResponse) GoString() string {
	return s.String()
}

func (s *InstanceHealerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstanceHealerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstanceHealerResponse) GetBody() *InstanceHealerResponseBody {
	return s.Body
}

func (s *InstanceHealerResponse) SetHeaders(v map[string]*string) *InstanceHealerResponse {
	s.Headers = v
	return s
}

func (s *InstanceHealerResponse) SetStatusCode(v int32) *InstanceHealerResponse {
	s.StatusCode = &v
	return s
}

func (s *InstanceHealerResponse) SetBody(v *InstanceHealerResponseBody) *InstanceHealerResponse {
	s.Body = v
	return s
}

func (s *InstanceHealerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
