// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadModuleResponse
	GetStatusCode() *int32
	SetBody(v *UploadModuleResponseBody) *UploadModuleResponse
	GetBody() *UploadModuleResponseBody
}

type UploadModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadModuleResponse) GoString() string {
	return s.String()
}

func (s *UploadModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadModuleResponse) GetBody() *UploadModuleResponseBody {
	return s.Body
}

func (s *UploadModuleResponse) SetHeaders(v map[string]*string) *UploadModuleResponse {
	s.Headers = v
	return s
}

func (s *UploadModuleResponse) SetStatusCode(v int32) *UploadModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadModuleResponse) SetBody(v *UploadModuleResponseBody) *UploadModuleResponse {
	s.Body = v
	return s
}

func (s *UploadModuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
