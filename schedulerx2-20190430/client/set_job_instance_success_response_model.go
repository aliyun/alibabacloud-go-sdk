// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetJobInstanceSuccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetJobInstanceSuccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetJobInstanceSuccessResponse
	GetStatusCode() *int32
	SetBody(v *SetJobInstanceSuccessResponseBody) *SetJobInstanceSuccessResponse
	GetBody() *SetJobInstanceSuccessResponseBody
}

type SetJobInstanceSuccessResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetJobInstanceSuccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetJobInstanceSuccessResponse) String() string {
	return dara.Prettify(s)
}

func (s SetJobInstanceSuccessResponse) GoString() string {
	return s.String()
}

func (s *SetJobInstanceSuccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetJobInstanceSuccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetJobInstanceSuccessResponse) GetBody() *SetJobInstanceSuccessResponseBody {
	return s.Body
}

func (s *SetJobInstanceSuccessResponse) SetHeaders(v map[string]*string) *SetJobInstanceSuccessResponse {
	s.Headers = v
	return s
}

func (s *SetJobInstanceSuccessResponse) SetStatusCode(v int32) *SetJobInstanceSuccessResponse {
	s.StatusCode = &v
	return s
}

func (s *SetJobInstanceSuccessResponse) SetBody(v *SetJobInstanceSuccessResponseBody) *SetJobInstanceSuccessResponse {
	s.Body = v
	return s
}

func (s *SetJobInstanceSuccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
