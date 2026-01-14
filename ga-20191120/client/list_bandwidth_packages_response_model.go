// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBandwidthPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBandwidthPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBandwidthPackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListBandwidthPackagesResponseBody) *ListBandwidthPackagesResponse
	GetBody() *ListBandwidthPackagesResponseBody
}

type ListBandwidthPackagesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBandwidthPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBandwidthPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBandwidthPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBandwidthPackagesResponse) GetBody() *ListBandwidthPackagesResponseBody {
	return s.Body
}

func (s *ListBandwidthPackagesResponse) SetHeaders(v map[string]*string) *ListBandwidthPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListBandwidthPackagesResponse) SetStatusCode(v int32) *ListBandwidthPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBandwidthPackagesResponse) SetBody(v *ListBandwidthPackagesResponseBody) *ListBandwidthPackagesResponse {
	s.Body = v
	return s
}

func (s *ListBandwidthPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
