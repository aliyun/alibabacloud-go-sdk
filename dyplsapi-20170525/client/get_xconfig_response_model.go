// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetXConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetXConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetXConfigResponseBody) *GetXConfigResponse
	GetBody() *GetXConfigResponseBody
}

type GetXConfigResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetXConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetXConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetXConfigResponse) GoString() string {
	return s.String()
}

func (s *GetXConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetXConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetXConfigResponse) GetBody() *GetXConfigResponseBody {
	return s.Body
}

func (s *GetXConfigResponse) SetHeaders(v map[string]*string) *GetXConfigResponse {
	s.Headers = v
	return s
}

func (s *GetXConfigResponse) SetStatusCode(v int32) *GetXConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetXConfigResponse) SetBody(v *GetXConfigResponseBody) *GetXConfigResponse {
	s.Body = v
	return s
}

func (s *GetXConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
