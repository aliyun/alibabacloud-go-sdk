// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigByNameResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigByNameResponseBody) *GetConfigByNameResponse
	GetBody() *GetConfigByNameResponseBody
}

type GetConfigByNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigByNameResponse) GoString() string {
	return s.String()
}

func (s *GetConfigByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigByNameResponse) GetBody() *GetConfigByNameResponseBody {
	return s.Body
}

func (s *GetConfigByNameResponse) SetHeaders(v map[string]*string) *GetConfigByNameResponse {
	s.Headers = v
	return s
}

func (s *GetConfigByNameResponse) SetStatusCode(v int32) *GetConfigByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigByNameResponse) SetBody(v *GetConfigByNameResponseBody) *GetConfigByNameResponse {
	s.Body = v
	return s
}

func (s *GetConfigByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
