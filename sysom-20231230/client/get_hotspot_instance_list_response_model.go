// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotInstanceListResponseBody) *GetHotspotInstanceListResponse
	GetBody() *GetHotspotInstanceListResponseBody
}

type GetHotspotInstanceListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotInstanceListResponse) GetBody() *GetHotspotInstanceListResponseBody {
	return s.Body
}

func (s *GetHotspotInstanceListResponse) SetHeaders(v map[string]*string) *GetHotspotInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotInstanceListResponse) SetStatusCode(v int32) *GetHotspotInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotInstanceListResponse) SetBody(v *GetHotspotInstanceListResponseBody) *GetHotspotInstanceListResponse {
	s.Body = v
	return s
}

func (s *GetHotspotInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
