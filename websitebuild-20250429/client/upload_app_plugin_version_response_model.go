// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAppPluginVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadAppPluginVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadAppPluginVersionResponse
	GetStatusCode() *int32
	SetBody(v *UploadAppPluginVersionResponseBody) *UploadAppPluginVersionResponse
	GetBody() *UploadAppPluginVersionResponseBody
}

type UploadAppPluginVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAppPluginVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAppPluginVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAppPluginVersionResponse) GoString() string {
	return s.String()
}

func (s *UploadAppPluginVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAppPluginVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadAppPluginVersionResponse) GetBody() *UploadAppPluginVersionResponseBody {
	return s.Body
}

func (s *UploadAppPluginVersionResponse) SetHeaders(v map[string]*string) *UploadAppPluginVersionResponse {
	s.Headers = v
	return s
}

func (s *UploadAppPluginVersionResponse) SetStatusCode(v int32) *UploadAppPluginVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAppPluginVersionResponse) SetBody(v *UploadAppPluginVersionResponseBody) *UploadAppPluginVersionResponse {
	s.Body = v
	return s
}

func (s *UploadAppPluginVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
