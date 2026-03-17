// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQosResponse
	GetStatusCode() *int32
	SetBody(v *CreateQosResponseBody) *CreateQosResponse
	GetBody() *CreateQosResponseBody
}

type CreateQosResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQosResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQosResponse) GoString() string {
	return s.String()
}

func (s *CreateQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQosResponse) GetBody() *CreateQosResponseBody {
	return s.Body
}

func (s *CreateQosResponse) SetHeaders(v map[string]*string) *CreateQosResponse {
	s.Headers = v
	return s
}

func (s *CreateQosResponse) SetStatusCode(v int32) *CreateQosResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQosResponse) SetBody(v *CreateQosResponseBody) *CreateQosResponse {
	s.Body = v
	return s
}

func (s *CreateQosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
