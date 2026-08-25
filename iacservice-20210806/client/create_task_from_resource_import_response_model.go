// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskFromResourceImportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTaskFromResourceImportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTaskFromResourceImportResponse
	GetStatusCode() *int32
	SetBody(v *CreateTaskFromResourceImportResponseBody) *CreateTaskFromResourceImportResponse
	GetBody() *CreateTaskFromResourceImportResponseBody
}

type CreateTaskFromResourceImportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskFromResourceImportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskFromResourceImportResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskFromResourceImportResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskFromResourceImportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTaskFromResourceImportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTaskFromResourceImportResponse) GetBody() *CreateTaskFromResourceImportResponseBody {
	return s.Body
}

func (s *CreateTaskFromResourceImportResponse) SetHeaders(v map[string]*string) *CreateTaskFromResourceImportResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskFromResourceImportResponse) SetStatusCode(v int32) *CreateTaskFromResourceImportResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskFromResourceImportResponse) SetBody(v *CreateTaskFromResourceImportResponseBody) *CreateTaskFromResourceImportResponse {
	s.Body = v
	return s
}

func (s *CreateTaskFromResourceImportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
