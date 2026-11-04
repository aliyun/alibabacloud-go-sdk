// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocalFileUploadJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLocalFileUploadJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLocalFileUploadJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateLocalFileUploadJobResponseBody) *CreateLocalFileUploadJobResponse
	GetBody() *CreateLocalFileUploadJobResponseBody
}

type CreateLocalFileUploadJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLocalFileUploadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLocalFileUploadJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLocalFileUploadJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLocalFileUploadJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLocalFileUploadJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLocalFileUploadJobResponse) GetBody() *CreateLocalFileUploadJobResponseBody {
	return s.Body
}

func (s *CreateLocalFileUploadJobResponse) SetHeaders(v map[string]*string) *CreateLocalFileUploadJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLocalFileUploadJobResponse) SetStatusCode(v int32) *CreateLocalFileUploadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLocalFileUploadJobResponse) SetBody(v *CreateLocalFileUploadJobResponseBody) *CreateLocalFileUploadJobResponse {
	s.Body = v
	return s
}

func (s *CreateLocalFileUploadJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
