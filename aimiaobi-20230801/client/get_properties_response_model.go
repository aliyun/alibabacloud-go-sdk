// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPropertiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPropertiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPropertiesResponse
	GetStatusCode() *int32
	SetBody(v *GetPropertiesResponseBody) *GetPropertiesResponse
	GetBody() *GetPropertiesResponseBody
}

type GetPropertiesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPropertiesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPropertiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPropertiesResponse) GetBody() *GetPropertiesResponseBody {
	return s.Body
}

func (s *GetPropertiesResponse) SetHeaders(v map[string]*string) *GetPropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetPropertiesResponse) SetStatusCode(v int32) *GetPropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPropertiesResponse) SetBody(v *GetPropertiesResponseBody) *GetPropertiesResponse {
	s.Body = v
	return s
}

func (s *GetPropertiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
