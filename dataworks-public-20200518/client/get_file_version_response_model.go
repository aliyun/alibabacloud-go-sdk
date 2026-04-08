// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFileVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFileVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetFileVersionResponseBody) *GetFileVersionResponse
	GetBody() *GetFileVersionResponseBody
}

type GetFileVersionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFileVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFileVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFileVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFileVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFileVersionResponse) GetBody() *GetFileVersionResponseBody {
	return s.Body
}

func (s *GetFileVersionResponse) SetHeaders(v map[string]*string) *GetFileVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFileVersionResponse) SetStatusCode(v int32) *GetFileVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileVersionResponse) SetBody(v *GetFileVersionResponseBody) *GetFileVersionResponse {
	s.Body = v
	return s
}

func (s *GetFileVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
