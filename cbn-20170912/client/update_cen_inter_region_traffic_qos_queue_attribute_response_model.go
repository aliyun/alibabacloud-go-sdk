// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCenInterRegionTrafficQosQueueAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCenInterRegionTrafficQosQueueAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCenInterRegionTrafficQosQueueAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) *UpdateCenInterRegionTrafficQosQueueAttributeResponse
	GetBody() *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody
}

type UpdateCenInterRegionTrafficQosQueueAttributeResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosQueueAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) GetBody() *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody {
	return s.Body
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) SetHeaders(v map[string]*string) *UpdateCenInterRegionTrafficQosQueueAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) SetStatusCode(v int32) *UpdateCenInterRegionTrafficQosQueueAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) SetBody(v *UpdateCenInterRegionTrafficQosQueueAttributeResponseBody) *UpdateCenInterRegionTrafficQosQueueAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosQueueAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
