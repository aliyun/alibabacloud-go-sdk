// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPlayKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppPlayKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppPlayKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetAppPlayKeyResponseBody) *GetAppPlayKeyResponse
	GetBody() *GetAppPlayKeyResponseBody
}

type GetAppPlayKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppPlayKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppPlayKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppPlayKeyResponse) GoString() string {
	return s.String()
}

func (s *GetAppPlayKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppPlayKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppPlayKeyResponse) GetBody() *GetAppPlayKeyResponseBody {
	return s.Body
}

func (s *GetAppPlayKeyResponse) SetHeaders(v map[string]*string) *GetAppPlayKeyResponse {
	s.Headers = v
	return s
}

func (s *GetAppPlayKeyResponse) SetStatusCode(v int32) *GetAppPlayKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppPlayKeyResponse) SetBody(v *GetAppPlayKeyResponseBody) *GetAppPlayKeyResponse {
	s.Body = v
	return s
}

func (s *GetAppPlayKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
