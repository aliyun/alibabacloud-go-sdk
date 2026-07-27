// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKgEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKgEntityResponse
	GetStatusCode() *int32
	SetBody(v *GetKgEntityResponseBody) *GetKgEntityResponse
	GetBody() *GetKgEntityResponseBody
}

type GetKgEntityResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKgEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKgEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKgEntityResponse) GoString() string {
	return s.String()
}

func (s *GetKgEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKgEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKgEntityResponse) GetBody() *GetKgEntityResponseBody {
	return s.Body
}

func (s *GetKgEntityResponse) SetHeaders(v map[string]*string) *GetKgEntityResponse {
	s.Headers = v
	return s
}

func (s *GetKgEntityResponse) SetStatusCode(v int32) *GetKgEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKgEntityResponse) SetBody(v *GetKgEntityResponseBody) *GetKgEntityResponse {
	s.Body = v
	return s
}

func (s *GetKgEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
