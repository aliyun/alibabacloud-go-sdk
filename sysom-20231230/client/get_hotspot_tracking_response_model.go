// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotTrackingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotTrackingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotTrackingResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotTrackingResponseBody) *GetHotspotTrackingResponse
	GetBody() *GetHotspotTrackingResponseBody
}

type GetHotspotTrackingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotTrackingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotTrackingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotTrackingResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotTrackingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotTrackingResponse) GetBody() *GetHotspotTrackingResponseBody {
	return s.Body
}

func (s *GetHotspotTrackingResponse) SetHeaders(v map[string]*string) *GetHotspotTrackingResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotTrackingResponse) SetStatusCode(v int32) *GetHotspotTrackingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotTrackingResponse) SetBody(v *GetHotspotTrackingResponseBody) *GetHotspotTrackingResponse {
	s.Body = v
	return s
}

func (s *GetHotspotTrackingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
