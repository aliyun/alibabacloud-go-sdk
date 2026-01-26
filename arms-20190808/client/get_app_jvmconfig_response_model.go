// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppJVMConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppJVMConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppJVMConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetAppJVMConfigResponseBody) *GetAppJVMConfigResponse
	GetBody() *GetAppJVMConfigResponseBody
}

type GetAppJVMConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppJVMConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppJVMConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppJVMConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAppJVMConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppJVMConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppJVMConfigResponse) GetBody() *GetAppJVMConfigResponseBody {
	return s.Body
}

func (s *GetAppJVMConfigResponse) SetHeaders(v map[string]*string) *GetAppJVMConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAppJVMConfigResponse) SetStatusCode(v int32) *GetAppJVMConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppJVMConfigResponse) SetBody(v *GetAppJVMConfigResponseBody) *GetAppJVMConfigResponse {
	s.Body = v
	return s
}

func (s *GetAppJVMConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
