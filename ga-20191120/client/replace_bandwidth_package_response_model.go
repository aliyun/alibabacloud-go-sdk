// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceBandwidthPackageResponseBody) *ReplaceBandwidthPackageResponse
	GetBody() *ReplaceBandwidthPackageResponseBody
}

type ReplaceBandwidthPackageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *ReplaceBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceBandwidthPackageResponse) GetBody() *ReplaceBandwidthPackageResponseBody {
	return s.Body
}

func (s *ReplaceBandwidthPackageResponse) SetHeaders(v map[string]*string) *ReplaceBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *ReplaceBandwidthPackageResponse) SetStatusCode(v int32) *ReplaceBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceBandwidthPackageResponse) SetBody(v *ReplaceBandwidthPackageResponseBody) *ReplaceBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *ReplaceBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
