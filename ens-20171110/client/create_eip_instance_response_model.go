// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEipInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEipInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEipInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateEipInstanceResponseBody) *CreateEipInstanceResponse
	GetBody() *CreateEipInstanceResponseBody
}

type CreateEipInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEipInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEipInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEipInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateEipInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEipInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEipInstanceResponse) GetBody() *CreateEipInstanceResponseBody {
	return s.Body
}

func (s *CreateEipInstanceResponse) SetHeaders(v map[string]*string) *CreateEipInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateEipInstanceResponse) SetStatusCode(v int32) *CreateEipInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEipInstanceResponse) SetBody(v *CreateEipInstanceResponseBody) *CreateEipInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateEipInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
