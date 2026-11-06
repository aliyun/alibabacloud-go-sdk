// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndBindAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndBindAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndBindAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndBindAppResponseBody) *CreateAndBindAppResponse
	GetBody() *CreateAndBindAppResponseBody
}

type CreateAndBindAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndBindAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndBindAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndBindAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAndBindAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndBindAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndBindAppResponse) GetBody() *CreateAndBindAppResponseBody {
	return s.Body
}

func (s *CreateAndBindAppResponse) SetHeaders(v map[string]*string) *CreateAndBindAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAndBindAppResponse) SetStatusCode(v int32) *CreateAndBindAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndBindAppResponse) SetBody(v *CreateAndBindAppResponseBody) *CreateAndBindAppResponse {
	s.Body = v
	return s
}

func (s *CreateAndBindAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
