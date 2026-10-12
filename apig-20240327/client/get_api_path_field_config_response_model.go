// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPathFieldConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiPathFieldConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiPathFieldConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetApiPathFieldConfigResponseBody) *GetApiPathFieldConfigResponse
	GetBody() *GetApiPathFieldConfigResponseBody
}

type GetApiPathFieldConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiPathFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiPathFieldConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiPathFieldConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiPathFieldConfigResponse) GetBody() *GetApiPathFieldConfigResponseBody {
	return s.Body
}

func (s *GetApiPathFieldConfigResponse) SetHeaders(v map[string]*string) *GetApiPathFieldConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApiPathFieldConfigResponse) SetStatusCode(v int32) *GetApiPathFieldConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiPathFieldConfigResponse) SetBody(v *GetApiPathFieldConfigResponseBody) *GetApiPathFieldConfigResponse {
	s.Body = v
	return s
}

func (s *GetApiPathFieldConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
