// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelVersionResponseBody) *CreateModelVersionResponse
	GetBody() *CreateModelVersionResponseBody
}

type CreateModelVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateModelVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelVersionResponse) GetBody() *CreateModelVersionResponseBody {
	return s.Body
}

func (s *CreateModelVersionResponse) SetHeaders(v map[string]*string) *CreateModelVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateModelVersionResponse) SetStatusCode(v int32) *CreateModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelVersionResponse) SetBody(v *CreateModelVersionResponseBody) *CreateModelVersionResponse {
	s.Body = v
	return s
}

func (s *CreateModelVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
