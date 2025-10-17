// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppLayoutResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppLayoutResponseBody) *CreateAppLayoutResponse
	GetBody() *CreateAppLayoutResponseBody
}

type CreateAppLayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLayoutResponse) GoString() string {
	return s.String()
}

func (s *CreateAppLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppLayoutResponse) GetBody() *CreateAppLayoutResponseBody {
	return s.Body
}

func (s *CreateAppLayoutResponse) SetHeaders(v map[string]*string) *CreateAppLayoutResponse {
	s.Headers = v
	return s
}

func (s *CreateAppLayoutResponse) SetStatusCode(v int32) *CreateAppLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppLayoutResponse) SetBody(v *CreateAppLayoutResponseBody) *CreateAppLayoutResponse {
	s.Body = v
	return s
}

func (s *CreateAppLayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
