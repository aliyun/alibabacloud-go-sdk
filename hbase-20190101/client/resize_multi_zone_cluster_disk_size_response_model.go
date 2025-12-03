// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeMultiZoneClusterDiskSizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeMultiZoneClusterDiskSizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeMultiZoneClusterDiskSizeResponse
	GetStatusCode() *int32
	SetBody(v *ResizeMultiZoneClusterDiskSizeResponseBody) *ResizeMultiZoneClusterDiskSizeResponse
	GetBody() *ResizeMultiZoneClusterDiskSizeResponseBody
}

type ResizeMultiZoneClusterDiskSizeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeMultiZoneClusterDiskSizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeMultiZoneClusterDiskSizeResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeMultiZoneClusterDiskSizeResponse) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) GetBody() *ResizeMultiZoneClusterDiskSizeResponseBody {
	return s.Body
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetHeaders(v map[string]*string) *ResizeMultiZoneClusterDiskSizeResponse {
	s.Headers = v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetStatusCode(v int32) *ResizeMultiZoneClusterDiskSizeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) SetBody(v *ResizeMultiZoneClusterDiskSizeResponseBody) *ResizeMultiZoneClusterDiskSizeResponse {
	s.Body = v
	return s
}

func (s *ResizeMultiZoneClusterDiskSizeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
