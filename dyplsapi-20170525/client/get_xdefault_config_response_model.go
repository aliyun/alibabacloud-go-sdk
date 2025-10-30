// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXDefaultConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetXDefaultConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetXDefaultConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetXDefaultConfigResponseBody) *GetXDefaultConfigResponse
	GetBody() *GetXDefaultConfigResponseBody
}

type GetXDefaultConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetXDefaultConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetXDefaultConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetXDefaultConfigResponse) GoString() string {
	return s.String()
}

func (s *GetXDefaultConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetXDefaultConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetXDefaultConfigResponse) GetBody() *GetXDefaultConfigResponseBody {
	return s.Body
}

func (s *GetXDefaultConfigResponse) SetHeaders(v map[string]*string) *GetXDefaultConfigResponse {
	s.Headers = v
	return s
}

func (s *GetXDefaultConfigResponse) SetStatusCode(v int32) *GetXDefaultConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetXDefaultConfigResponse) SetBody(v *GetXDefaultConfigResponseBody) *GetXDefaultConfigResponse {
	s.Body = v
	return s
}

func (s *GetXDefaultConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
