// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocalFileUploadJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLocalFileUploadJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLocalFileUploadJobResponse
	GetStatusCode() *int32
	SetBody(v *QueryLocalFileUploadJobResponseBody) *QueryLocalFileUploadJobResponse
	GetBody() *QueryLocalFileUploadJobResponseBody
}

type QueryLocalFileUploadJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLocalFileUploadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLocalFileUploadJobResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalFileUploadJobResponse) GoString() string {
	return s.String()
}

func (s *QueryLocalFileUploadJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLocalFileUploadJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLocalFileUploadJobResponse) GetBody() *QueryLocalFileUploadJobResponseBody {
	return s.Body
}

func (s *QueryLocalFileUploadJobResponse) SetHeaders(v map[string]*string) *QueryLocalFileUploadJobResponse {
	s.Headers = v
	return s
}

func (s *QueryLocalFileUploadJobResponse) SetStatusCode(v int32) *QueryLocalFileUploadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLocalFileUploadJobResponse) SetBody(v *QueryLocalFileUploadJobResponseBody) *QueryLocalFileUploadJobResponse {
	s.Body = v
	return s
}

func (s *QueryLocalFileUploadJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
