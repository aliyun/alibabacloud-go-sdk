// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenInterRegionTrafficQosQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenInterRegionTrafficQosQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenInterRegionTrafficQosQueueResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenInterRegionTrafficQosQueueResponseBody) *CreateCenInterRegionTrafficQosQueueResponse
	GetBody() *CreateCenInterRegionTrafficQosQueueResponseBody
}

type CreateCenInterRegionTrafficQosQueueResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenInterRegionTrafficQosQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenInterRegionTrafficQosQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenInterRegionTrafficQosQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenInterRegionTrafficQosQueueResponse) GetBody() *CreateCenInterRegionTrafficQosQueueResponseBody {
	return s.Body
}

func (s *CreateCenInterRegionTrafficQosQueueResponse) SetHeaders(v map[string]*string) *CreateCenInterRegionTrafficQosQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueResponse) SetStatusCode(v int32) *CreateCenInterRegionTrafficQosQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueResponse) SetBody(v *CreateCenInterRegionTrafficQosQueueResponseBody) *CreateCenInterRegionTrafficQosQueueResponse {
	s.Body = v
	return s
}

func (s *CreateCenInterRegionTrafficQosQueueResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
