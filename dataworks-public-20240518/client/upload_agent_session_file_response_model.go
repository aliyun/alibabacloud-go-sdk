// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSessionFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadAgentSessionFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadAgentSessionFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadAgentSessionFileResponseBody) *UploadAgentSessionFileResponse
	GetBody() *UploadAgentSessionFileResponseBody
}

type UploadAgentSessionFileResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAgentSessionFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAgentSessionFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileResponse) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAgentSessionFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadAgentSessionFileResponse) GetBody() *UploadAgentSessionFileResponseBody {
	return s.Body
}

func (s *UploadAgentSessionFileResponse) SetHeaders(v map[string]*string) *UploadAgentSessionFileResponse {
	s.Headers = v
	return s
}

func (s *UploadAgentSessionFileResponse) SetStatusCode(v int32) *UploadAgentSessionFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAgentSessionFileResponse) SetBody(v *UploadAgentSessionFileResponseBody) *UploadAgentSessionFileResponse {
	s.Body = v
	return s
}

func (s *UploadAgentSessionFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
