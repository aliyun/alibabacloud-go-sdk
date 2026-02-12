// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsRegionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsRegionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsRegionListResponse
	GetStatusCode() *int32
	SetBody(v *OnsRegionListResponseBody) *OnsRegionListResponse
	GetBody() *OnsRegionListResponseBody
}

type OnsRegionListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsRegionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsRegionListResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsRegionListResponse) GoString() string {
	return s.String()
}

func (s *OnsRegionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsRegionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsRegionListResponse) GetBody() *OnsRegionListResponseBody {
	return s.Body
}

func (s *OnsRegionListResponse) SetHeaders(v map[string]*string) *OnsRegionListResponse {
	s.Headers = v
	return s
}

func (s *OnsRegionListResponse) SetStatusCode(v int32) *OnsRegionListResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsRegionListResponse) SetBody(v *OnsRegionListResponseBody) *OnsRegionListResponse {
	s.Body = v
	return s
}

func (s *OnsRegionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
