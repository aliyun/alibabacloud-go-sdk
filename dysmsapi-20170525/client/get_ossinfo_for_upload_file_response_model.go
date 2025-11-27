// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOSSInfoForUploadFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOSSInfoForUploadFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOSSInfoForUploadFileResponse
	GetStatusCode() *int32
	SetBody(v *GetOSSInfoForUploadFileResponseBody) *GetOSSInfoForUploadFileResponse
	GetBody() *GetOSSInfoForUploadFileResponseBody
}

type GetOSSInfoForUploadFileResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOSSInfoForUploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOSSInfoForUploadFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOSSInfoForUploadFileResponse) GoString() string {
	return s.String()
}

func (s *GetOSSInfoForUploadFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOSSInfoForUploadFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOSSInfoForUploadFileResponse) GetBody() *GetOSSInfoForUploadFileResponseBody {
	return s.Body
}

func (s *GetOSSInfoForUploadFileResponse) SetHeaders(v map[string]*string) *GetOSSInfoForUploadFileResponse {
	s.Headers = v
	return s
}

func (s *GetOSSInfoForUploadFileResponse) SetStatusCode(v int32) *GetOSSInfoForUploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOSSInfoForUploadFileResponse) SetBody(v *GetOSSInfoForUploadFileResponseBody) *GetOSSInfoForUploadFileResponse {
	s.Body = v
	return s
}

func (s *GetOSSInfoForUploadFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
