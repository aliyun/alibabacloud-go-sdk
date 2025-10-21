// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavepointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSavepointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSavepointResponse
	GetStatusCode() *int32
	SetBody(v *CreateSavepointResponseBody) *CreateSavepointResponse
	GetBody() *CreateSavepointResponseBody
}

type CreateSavepointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSavepointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSavepointResponse) GoString() string {
	return s.String()
}

func (s *CreateSavepointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSavepointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSavepointResponse) GetBody() *CreateSavepointResponseBody {
	return s.Body
}

func (s *CreateSavepointResponse) SetHeaders(v map[string]*string) *CreateSavepointResponse {
	s.Headers = v
	return s
}

func (s *CreateSavepointResponse) SetStatusCode(v int32) *CreateSavepointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSavepointResponse) SetBody(v *CreateSavepointResponseBody) *CreateSavepointResponse {
	s.Body = v
	return s
}

func (s *CreateSavepointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
