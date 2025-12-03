// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomFieldOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomFieldOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomFieldOptionResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomFieldOptionResponseBody) *GetCustomFieldOptionResponse
	GetBody() *GetCustomFieldOptionResponseBody
}

type GetCustomFieldOptionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomFieldOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomFieldOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldOptionResponse) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomFieldOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomFieldOptionResponse) GetBody() *GetCustomFieldOptionResponseBody {
	return s.Body
}

func (s *GetCustomFieldOptionResponse) SetHeaders(v map[string]*string) *GetCustomFieldOptionResponse {
	s.Headers = v
	return s
}

func (s *GetCustomFieldOptionResponse) SetStatusCode(v int32) *GetCustomFieldOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomFieldOptionResponse) SetBody(v *GetCustomFieldOptionResponseBody) *GetCustomFieldOptionResponse {
	s.Body = v
	return s
}

func (s *GetCustomFieldOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
