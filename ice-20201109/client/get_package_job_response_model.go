// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackageJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPackageJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPackageJobResponse
	GetStatusCode() *int32
	SetBody(v *GetPackageJobResponseBody) *GetPackageJobResponse
	GetBody() *GetPackageJobResponseBody
}

type GetPackageJobResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPackageJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPackageJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPackageJobResponse) GoString() string {
	return s.String()
}

func (s *GetPackageJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPackageJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPackageJobResponse) GetBody() *GetPackageJobResponseBody {
	return s.Body
}

func (s *GetPackageJobResponse) SetHeaders(v map[string]*string) *GetPackageJobResponse {
	s.Headers = v
	return s
}

func (s *GetPackageJobResponse) SetStatusCode(v int32) *GetPackageJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPackageJobResponse) SetBody(v *GetPackageJobResponseBody) *GetPackageJobResponse {
	s.Body = v
	return s
}

func (s *GetPackageJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
