// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataProcessJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataProcessJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataProcessJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataProcessJobResponseBody) *CreateDataProcessJobResponse
	GetBody() *CreateDataProcessJobResponseBody
}

type CreateDataProcessJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataProcessJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataProcessJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataProcessJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDataProcessJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataProcessJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataProcessJobResponse) GetBody() *CreateDataProcessJobResponseBody {
	return s.Body
}

func (s *CreateDataProcessJobResponse) SetHeaders(v map[string]*string) *CreateDataProcessJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDataProcessJobResponse) SetStatusCode(v int32) *CreateDataProcessJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataProcessJobResponse) SetBody(v *CreateDataProcessJobResponseBody) *CreateDataProcessJobResponse {
	s.Body = v
	return s
}

func (s *CreateDataProcessJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
