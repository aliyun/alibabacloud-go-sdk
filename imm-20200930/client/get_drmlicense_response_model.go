// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDRMLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDRMLicenseResponse
	GetStatusCode() *int32
	SetBody(v *GetDRMLicenseResponseBody) *GetDRMLicenseResponse
	GetBody() *GetDRMLicenseResponseBody
}

type GetDRMLicenseResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDRMLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDRMLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDRMLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDRMLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDRMLicenseResponse) GetBody() *GetDRMLicenseResponseBody {
	return s.Body
}

func (s *GetDRMLicenseResponse) SetHeaders(v map[string]*string) *GetDRMLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetDRMLicenseResponse) SetStatusCode(v int32) *GetDRMLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDRMLicenseResponse) SetBody(v *GetDRMLicenseResponseBody) *GetDRMLicenseResponse {
	s.Body = v
	return s
}

func (s *GetDRMLicenseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
