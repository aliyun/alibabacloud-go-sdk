// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSwitchRegionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSwitchRegionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSwitchRegionDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSwitchRegionDetailResponseBody) *GetSwitchRegionDetailResponse
	GetBody() *GetSwitchRegionDetailResponseBody
}

type GetSwitchRegionDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSwitchRegionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSwitchRegionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSwitchRegionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSwitchRegionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSwitchRegionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSwitchRegionDetailResponse) GetBody() *GetSwitchRegionDetailResponseBody {
	return s.Body
}

func (s *GetSwitchRegionDetailResponse) SetHeaders(v map[string]*string) *GetSwitchRegionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSwitchRegionDetailResponse) SetStatusCode(v int32) *GetSwitchRegionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSwitchRegionDetailResponse) SetBody(v *GetSwitchRegionDetailResponseBody) *GetSwitchRegionDetailResponse {
	s.Body = v
	return s
}

func (s *GetSwitchRegionDetailResponse) Validate() error {
	return dara.Validate(s)
}
