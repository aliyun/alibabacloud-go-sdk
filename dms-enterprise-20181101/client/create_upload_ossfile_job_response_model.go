// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadOSSFileJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadOSSFileJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadOSSFileJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadOSSFileJobResponseBody) *CreateUploadOSSFileJobResponse
	GetBody() *CreateUploadOSSFileJobResponseBody
}

type CreateUploadOSSFileJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadOSSFileJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadOSSFileJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadOSSFileJobResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadOSSFileJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadOSSFileJobResponse) GetBody() *CreateUploadOSSFileJobResponseBody {
	return s.Body
}

func (s *CreateUploadOSSFileJobResponse) SetHeaders(v map[string]*string) *CreateUploadOSSFileJobResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadOSSFileJobResponse) SetStatusCode(v int32) *CreateUploadOSSFileJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadOSSFileJobResponse) SetBody(v *CreateUploadOSSFileJobResponseBody) *CreateUploadOSSFileJobResponse {
	s.Body = v
	return s
}

func (s *CreateUploadOSSFileJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
