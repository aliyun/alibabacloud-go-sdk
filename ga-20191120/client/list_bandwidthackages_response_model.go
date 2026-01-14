// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBandwidthackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBandwidthackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBandwidthackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListBandwidthackagesResponseBody) *ListBandwidthackagesResponse
	GetBody() *ListBandwidthackagesResponseBody
}

type ListBandwidthackagesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBandwidthackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBandwidthackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthackagesResponse) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBandwidthackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBandwidthackagesResponse) GetBody() *ListBandwidthackagesResponseBody {
	return s.Body
}

func (s *ListBandwidthackagesResponse) SetHeaders(v map[string]*string) *ListBandwidthackagesResponse {
	s.Headers = v
	return s
}

func (s *ListBandwidthackagesResponse) SetStatusCode(v int32) *ListBandwidthackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBandwidthackagesResponse) SetBody(v *ListBandwidthackagesResponseBody) *ListBandwidthackagesResponse {
	s.Body = v
	return s
}

func (s *ListBandwidthackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
