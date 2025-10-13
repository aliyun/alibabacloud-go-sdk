// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradeApplicationApmServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DowngradeApplicationApmServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DowngradeApplicationApmServiceResponse
	GetStatusCode() *int32
	SetBody(v *DowngradeApplicationApmServiceResponseBody) *DowngradeApplicationApmServiceResponse
	GetBody() *DowngradeApplicationApmServiceResponseBody
}

type DowngradeApplicationApmServiceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DowngradeApplicationApmServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DowngradeApplicationApmServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DowngradeApplicationApmServiceResponse) GoString() string {
	return s.String()
}

func (s *DowngradeApplicationApmServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DowngradeApplicationApmServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DowngradeApplicationApmServiceResponse) GetBody() *DowngradeApplicationApmServiceResponseBody {
	return s.Body
}

func (s *DowngradeApplicationApmServiceResponse) SetHeaders(v map[string]*string) *DowngradeApplicationApmServiceResponse {
	s.Headers = v
	return s
}

func (s *DowngradeApplicationApmServiceResponse) SetStatusCode(v int32) *DowngradeApplicationApmServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DowngradeApplicationApmServiceResponse) SetBody(v *DowngradeApplicationApmServiceResponseBody) *DowngradeApplicationApmServiceResponse {
	s.Body = v
	return s
}

func (s *DowngradeApplicationApmServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
