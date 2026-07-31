// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSemanticFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadSemanticFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadSemanticFileResponse
	GetStatusCode() *int32
	SetBody(v *UploadSemanticFileResponseBody) *UploadSemanticFileResponse
	GetBody() *UploadSemanticFileResponseBody
}

type UploadSemanticFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadSemanticFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadSemanticFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadSemanticFileResponse) GoString() string {
	return s.String()
}

func (s *UploadSemanticFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadSemanticFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadSemanticFileResponse) GetBody() *UploadSemanticFileResponseBody {
	return s.Body
}

func (s *UploadSemanticFileResponse) SetHeaders(v map[string]*string) *UploadSemanticFileResponse {
	s.Headers = v
	return s
}

func (s *UploadSemanticFileResponse) SetStatusCode(v int32) *UploadSemanticFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadSemanticFileResponse) SetBody(v *UploadSemanticFileResponseBody) *UploadSemanticFileResponse {
	s.Body = v
	return s
}

func (s *UploadSemanticFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
