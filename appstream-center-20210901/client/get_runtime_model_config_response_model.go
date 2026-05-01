// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeModelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuntimeModelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuntimeModelConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetRuntimeModelConfigResponseBody) *GetRuntimeModelConfigResponse
	GetBody() *GetRuntimeModelConfigResponseBody
}

type GetRuntimeModelConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuntimeModelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuntimeModelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeModelConfigResponse) GoString() string {
	return s.String()
}

func (s *GetRuntimeModelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuntimeModelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuntimeModelConfigResponse) GetBody() *GetRuntimeModelConfigResponseBody {
	return s.Body
}

func (s *GetRuntimeModelConfigResponse) SetHeaders(v map[string]*string) *GetRuntimeModelConfigResponse {
	s.Headers = v
	return s
}

func (s *GetRuntimeModelConfigResponse) SetStatusCode(v int32) *GetRuntimeModelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuntimeModelConfigResponse) SetBody(v *GetRuntimeModelConfigResponseBody) *GetRuntimeModelConfigResponse {
	s.Body = v
	return s
}

func (s *GetRuntimeModelConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
