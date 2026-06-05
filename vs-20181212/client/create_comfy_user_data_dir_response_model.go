// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComfyUserDataDirResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateComfyUserDataDirResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateComfyUserDataDirResponse
	GetStatusCode() *int32
	SetBody(v *CreateComfyUserDataDirResponseBody) *CreateComfyUserDataDirResponse
	GetBody() *CreateComfyUserDataDirResponseBody
}

type CreateComfyUserDataDirResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateComfyUserDataDirResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateComfyUserDataDirResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateComfyUserDataDirResponse) GoString() string {
	return s.String()
}

func (s *CreateComfyUserDataDirResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateComfyUserDataDirResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateComfyUserDataDirResponse) GetBody() *CreateComfyUserDataDirResponseBody {
	return s.Body
}

func (s *CreateComfyUserDataDirResponse) SetHeaders(v map[string]*string) *CreateComfyUserDataDirResponse {
	s.Headers = v
	return s
}

func (s *CreateComfyUserDataDirResponse) SetStatusCode(v int32) *CreateComfyUserDataDirResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateComfyUserDataDirResponse) SetBody(v *CreateComfyUserDataDirResponseBody) *CreateComfyUserDataDirResponse {
	s.Body = v
	return s
}

func (s *CreateComfyUserDataDirResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
