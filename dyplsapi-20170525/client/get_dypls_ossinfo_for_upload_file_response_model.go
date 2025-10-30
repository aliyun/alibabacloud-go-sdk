// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDyplsOSSInfoForUploadFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDyplsOSSInfoForUploadFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDyplsOSSInfoForUploadFileResponse
	GetStatusCode() *int32
	SetBody(v *GetDyplsOSSInfoForUploadFileResponseBody) *GetDyplsOSSInfoForUploadFileResponse
	GetBody() *GetDyplsOSSInfoForUploadFileResponseBody
}

type GetDyplsOSSInfoForUploadFileResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDyplsOSSInfoForUploadFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDyplsOSSInfoForUploadFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDyplsOSSInfoForUploadFileResponse) GoString() string {
	return s.String()
}

func (s *GetDyplsOSSInfoForUploadFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDyplsOSSInfoForUploadFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDyplsOSSInfoForUploadFileResponse) GetBody() *GetDyplsOSSInfoForUploadFileResponseBody {
	return s.Body
}

func (s *GetDyplsOSSInfoForUploadFileResponse) SetHeaders(v map[string]*string) *GetDyplsOSSInfoForUploadFileResponse {
	s.Headers = v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponse) SetStatusCode(v int32) *GetDyplsOSSInfoForUploadFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponse) SetBody(v *GetDyplsOSSInfoForUploadFileResponseBody) *GetDyplsOSSInfoForUploadFileResponse {
	s.Body = v
	return s
}

func (s *GetDyplsOSSInfoForUploadFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
