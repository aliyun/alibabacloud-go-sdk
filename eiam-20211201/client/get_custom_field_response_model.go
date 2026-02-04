// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomFieldResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomFieldResponseBody) *GetCustomFieldResponse
	GetBody() *GetCustomFieldResponseBody
}

type GetCustomFieldResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldResponse) GoString() string {
	return s.String()
}

func (s *GetCustomFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomFieldResponse) GetBody() *GetCustomFieldResponseBody {
	return s.Body
}

func (s *GetCustomFieldResponse) SetHeaders(v map[string]*string) *GetCustomFieldResponse {
	s.Headers = v
	return s
}

func (s *GetCustomFieldResponse) SetStatusCode(v int32) *GetCustomFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomFieldResponse) SetBody(v *GetCustomFieldResponseBody) *GetCustomFieldResponse {
	s.Body = v
	return s
}

func (s *GetCustomFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
