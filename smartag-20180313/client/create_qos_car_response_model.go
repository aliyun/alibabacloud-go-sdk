// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosCarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQosCarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQosCarResponse
	GetStatusCode() *int32
	SetBody(v *CreateQosCarResponseBody) *CreateQosCarResponse
	GetBody() *CreateQosCarResponseBody
}

type CreateQosCarResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQosCarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQosCarResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQosCarResponse) GoString() string {
	return s.String()
}

func (s *CreateQosCarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQosCarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQosCarResponse) GetBody() *CreateQosCarResponseBody {
	return s.Body
}

func (s *CreateQosCarResponse) SetHeaders(v map[string]*string) *CreateQosCarResponse {
	s.Headers = v
	return s
}

func (s *CreateQosCarResponse) SetStatusCode(v int32) *CreateQosCarResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQosCarResponse) SetBody(v *CreateQosCarResponseBody) *CreateQosCarResponse {
	s.Body = v
	return s
}

func (s *CreateQosCarResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
