// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJavaStartUpConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJavaStartUpConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJavaStartUpConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetJavaStartUpConfigResponseBody) *GetJavaStartUpConfigResponse
	GetBody() *GetJavaStartUpConfigResponseBody
}

type GetJavaStartUpConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJavaStartUpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJavaStartUpConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJavaStartUpConfigResponse) GoString() string {
	return s.String()
}

func (s *GetJavaStartUpConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJavaStartUpConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJavaStartUpConfigResponse) GetBody() *GetJavaStartUpConfigResponseBody {
	return s.Body
}

func (s *GetJavaStartUpConfigResponse) SetHeaders(v map[string]*string) *GetJavaStartUpConfigResponse {
	s.Headers = v
	return s
}

func (s *GetJavaStartUpConfigResponse) SetStatusCode(v int32) *GetJavaStartUpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJavaStartUpConfigResponse) SetBody(v *GetJavaStartUpConfigResponseBody) *GetJavaStartUpConfigResponse {
	s.Body = v
	return s
}

func (s *GetJavaStartUpConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
