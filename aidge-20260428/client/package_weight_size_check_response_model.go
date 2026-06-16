// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackageWeightSizeCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PackageWeightSizeCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PackageWeightSizeCheckResponse
	GetStatusCode() *int32
	SetBody(v *PackageWeightSizeCheckResponseBody) *PackageWeightSizeCheckResponse
	GetBody() *PackageWeightSizeCheckResponseBody
}

type PackageWeightSizeCheckResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PackageWeightSizeCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PackageWeightSizeCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s PackageWeightSizeCheckResponse) GoString() string {
	return s.String()
}

func (s *PackageWeightSizeCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PackageWeightSizeCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PackageWeightSizeCheckResponse) GetBody() *PackageWeightSizeCheckResponseBody {
	return s.Body
}

func (s *PackageWeightSizeCheckResponse) SetHeaders(v map[string]*string) *PackageWeightSizeCheckResponse {
	s.Headers = v
	return s
}

func (s *PackageWeightSizeCheckResponse) SetStatusCode(v int32) *PackageWeightSizeCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *PackageWeightSizeCheckResponse) SetBody(v *PackageWeightSizeCheckResponseBody) *PackageWeightSizeCheckResponse {
	s.Body = v
	return s
}

func (s *PackageWeightSizeCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
