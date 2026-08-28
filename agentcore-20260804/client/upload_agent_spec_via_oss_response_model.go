// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSpecViaOssResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadAgentSpecViaOssResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadAgentSpecViaOssResponse
	GetStatusCode() *int32
	SetBody(v *UploadAgentSpecViaOssResponseBody) *UploadAgentSpecViaOssResponse
	GetBody() *UploadAgentSpecViaOssResponseBody
}

type UploadAgentSpecViaOssResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAgentSpecViaOssResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAgentSpecViaOssResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSpecViaOssResponse) GoString() string {
	return s.String()
}

func (s *UploadAgentSpecViaOssResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAgentSpecViaOssResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadAgentSpecViaOssResponse) GetBody() *UploadAgentSpecViaOssResponseBody {
	return s.Body
}

func (s *UploadAgentSpecViaOssResponse) SetHeaders(v map[string]*string) *UploadAgentSpecViaOssResponse {
	s.Headers = v
	return s
}

func (s *UploadAgentSpecViaOssResponse) SetStatusCode(v int32) *UploadAgentSpecViaOssResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAgentSpecViaOssResponse) SetBody(v *UploadAgentSpecViaOssResponseBody) *UploadAgentSpecViaOssResponse {
	s.Body = v
	return s
}

func (s *UploadAgentSpecViaOssResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
