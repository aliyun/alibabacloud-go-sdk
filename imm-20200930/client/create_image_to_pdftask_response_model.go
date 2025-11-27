// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageToPDFTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImageToPDFTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImageToPDFTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateImageToPDFTaskResponseBody) *CreateImageToPDFTaskResponse
	GetBody() *CreateImageToPDFTaskResponseBody
}

type CreateImageToPDFTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImageToPDFTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImageToPDFTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImageToPDFTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImageToPDFTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImageToPDFTaskResponse) GetBody() *CreateImageToPDFTaskResponseBody {
	return s.Body
}

func (s *CreateImageToPDFTaskResponse) SetHeaders(v map[string]*string) *CreateImageToPDFTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateImageToPDFTaskResponse) SetStatusCode(v int32) *CreateImageToPDFTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageToPDFTaskResponse) SetBody(v *CreateImageToPDFTaskResponseBody) *CreateImageToPDFTaskResponse {
	s.Body = v
	return s
}

func (s *CreateImageToPDFTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
