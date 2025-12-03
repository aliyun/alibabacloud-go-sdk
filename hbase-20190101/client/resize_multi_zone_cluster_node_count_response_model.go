// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeMultiZoneClusterNodeCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResizeMultiZoneClusterNodeCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResizeMultiZoneClusterNodeCountResponse
	GetStatusCode() *int32
	SetBody(v *ResizeMultiZoneClusterNodeCountResponseBody) *ResizeMultiZoneClusterNodeCountResponse
	GetBody() *ResizeMultiZoneClusterNodeCountResponseBody
}

type ResizeMultiZoneClusterNodeCountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResizeMultiZoneClusterNodeCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResizeMultiZoneClusterNodeCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ResizeMultiZoneClusterNodeCountResponse) GoString() string {
	return s.String()
}

func (s *ResizeMultiZoneClusterNodeCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResizeMultiZoneClusterNodeCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResizeMultiZoneClusterNodeCountResponse) GetBody() *ResizeMultiZoneClusterNodeCountResponseBody {
	return s.Body
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetHeaders(v map[string]*string) *ResizeMultiZoneClusterNodeCountResponse {
	s.Headers = v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetStatusCode(v int32) *ResizeMultiZoneClusterNodeCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponse) SetBody(v *ResizeMultiZoneClusterNodeCountResponseBody) *ResizeMultiZoneClusterNodeCountResponse {
	s.Body = v
	return s
}

func (s *ResizeMultiZoneClusterNodeCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
