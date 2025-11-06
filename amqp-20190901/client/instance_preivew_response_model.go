// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstancePreivewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstancePreivewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstancePreivewResponse
	GetStatusCode() *int32
	SetBody(v *InstancePreivewResponseBody) *InstancePreivewResponse
	GetBody() *InstancePreivewResponseBody
}

type InstancePreivewResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstancePreivewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstancePreivewResponse) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewResponse) GoString() string {
	return s.String()
}

func (s *InstancePreivewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstancePreivewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstancePreivewResponse) GetBody() *InstancePreivewResponseBody {
	return s.Body
}

func (s *InstancePreivewResponse) SetHeaders(v map[string]*string) *InstancePreivewResponse {
	s.Headers = v
	return s
}

func (s *InstancePreivewResponse) SetStatusCode(v int32) *InstancePreivewResponse {
	s.StatusCode = &v
	return s
}

func (s *InstancePreivewResponse) SetBody(v *InstancePreivewResponseBody) *InstancePreivewResponse {
	s.Body = v
	return s
}

func (s *InstancePreivewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
