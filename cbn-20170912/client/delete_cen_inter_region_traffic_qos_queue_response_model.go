// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenInterRegionTrafficQosQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenInterRegionTrafficQosQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenInterRegionTrafficQosQueueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenInterRegionTrafficQosQueueResponseBody) *DeleteCenInterRegionTrafficQosQueueResponse
	GetBody() *DeleteCenInterRegionTrafficQosQueueResponseBody
}

type DeleteCenInterRegionTrafficQosQueueResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenInterRegionTrafficQosQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) GetBody() *DeleteCenInterRegionTrafficQosQueueResponseBody {
	return s.Body
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) SetHeaders(v map[string]*string) *DeleteCenInterRegionTrafficQosQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) SetStatusCode(v int32) *DeleteCenInterRegionTrafficQosQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) SetBody(v *DeleteCenInterRegionTrafficQosQueueResponseBody) *DeleteCenInterRegionTrafficQosQueueResponse {
	s.Body = v
	return s
}

func (s *DeleteCenInterRegionTrafficQosQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
