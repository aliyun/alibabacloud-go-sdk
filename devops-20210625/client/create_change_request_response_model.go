// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChangeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChangeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChangeRequestResponse
	GetStatusCode() *int32
	SetBody(v *CreateChangeRequestResponseBody) *CreateChangeRequestResponse
	GetBody() *CreateChangeRequestResponseBody
}

type CreateChangeRequestResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChangeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChangeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateChangeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChangeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChangeRequestResponse) GetBody() *CreateChangeRequestResponseBody {
	return s.Body
}

func (s *CreateChangeRequestResponse) SetHeaders(v map[string]*string) *CreateChangeRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateChangeRequestResponse) SetStatusCode(v int32) *CreateChangeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChangeRequestResponse) SetBody(v *CreateChangeRequestResponseBody) *CreateChangeRequestResponse {
	s.Body = v
	return s
}

func (s *CreateChangeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
