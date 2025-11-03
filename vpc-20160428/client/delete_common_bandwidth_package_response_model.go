// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommonBandwidthPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCommonBandwidthPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCommonBandwidthPackageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCommonBandwidthPackageResponseBody) *DeleteCommonBandwidthPackageResponse
	GetBody() *DeleteCommonBandwidthPackageResponseBody
}

type DeleteCommonBandwidthPackageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCommonBandwidthPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCommonBandwidthPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommonBandwidthPackageResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommonBandwidthPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCommonBandwidthPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCommonBandwidthPackageResponse) GetBody() *DeleteCommonBandwidthPackageResponseBody {
	return s.Body
}

func (s *DeleteCommonBandwidthPackageResponse) SetHeaders(v map[string]*string) *DeleteCommonBandwidthPackageResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommonBandwidthPackageResponse) SetStatusCode(v int32) *DeleteCommonBandwidthPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommonBandwidthPackageResponse) SetBody(v *DeleteCommonBandwidthPackageResponseBody) *DeleteCommonBandwidthPackageResponse {
	s.Body = v
	return s
}

func (s *DeleteCommonBandwidthPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
