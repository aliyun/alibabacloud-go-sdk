// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUrlUploadJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelUrlUploadJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelUrlUploadJobsResponse
	GetStatusCode() *int32
	SetBody(v *CancelUrlUploadJobsResponseBody) *CancelUrlUploadJobsResponse
	GetBody() *CancelUrlUploadJobsResponseBody
}

type CancelUrlUploadJobsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelUrlUploadJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelUrlUploadJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelUrlUploadJobsResponse) GoString() string {
	return s.String()
}

func (s *CancelUrlUploadJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelUrlUploadJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelUrlUploadJobsResponse) GetBody() *CancelUrlUploadJobsResponseBody {
	return s.Body
}

func (s *CancelUrlUploadJobsResponse) SetHeaders(v map[string]*string) *CancelUrlUploadJobsResponse {
	s.Headers = v
	return s
}

func (s *CancelUrlUploadJobsResponse) SetStatusCode(v int32) *CancelUrlUploadJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelUrlUploadJobsResponse) SetBody(v *CancelUrlUploadJobsResponseBody) *CancelUrlUploadJobsResponse {
	s.Body = v
	return s
}

func (s *CancelUrlUploadJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
