// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLocalFileUploadJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLocalFileUploadJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLocalFileUploadJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLocalFileUploadJobResponseBody) *DeleteLocalFileUploadJobResponse
	GetBody() *DeleteLocalFileUploadJobResponseBody
}

type DeleteLocalFileUploadJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLocalFileUploadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLocalFileUploadJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLocalFileUploadJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteLocalFileUploadJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLocalFileUploadJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLocalFileUploadJobResponse) GetBody() *DeleteLocalFileUploadJobResponseBody {
	return s.Body
}

func (s *DeleteLocalFileUploadJobResponse) SetHeaders(v map[string]*string) *DeleteLocalFileUploadJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteLocalFileUploadJobResponse) SetStatusCode(v int32) *DeleteLocalFileUploadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLocalFileUploadJobResponse) SetBody(v *DeleteLocalFileUploadJobResponseBody) *DeleteLocalFileUploadJobResponse {
	s.Body = v
	return s
}

func (s *DeleteLocalFileUploadJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
