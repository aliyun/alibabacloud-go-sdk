// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPackageJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitPackageJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitPackageJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitPackageJobResponseBody) *SubmitPackageJobResponse
	GetBody() *SubmitPackageJobResponseBody
}

type SubmitPackageJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitPackageJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitPackageJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitPackageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitPackageJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitPackageJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitPackageJobResponse) GetBody() *SubmitPackageJobResponseBody {
	return s.Body
}

func (s *SubmitPackageJobResponse) SetHeaders(v map[string]*string) *SubmitPackageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitPackageJobResponse) SetStatusCode(v int32) *SubmitPackageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitPackageJobResponse) SetBody(v *SubmitPackageJobResponseBody) *SubmitPackageJobResponse {
	s.Body = v
	return s
}

func (s *SubmitPackageJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
