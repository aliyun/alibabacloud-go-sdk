// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotPidListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotspotPidListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotspotPidListResponse
	GetStatusCode() *int32
	SetBody(v *GetHotspotPidListResponseBody) *GetHotspotPidListResponse
	GetBody() *GetHotspotPidListResponseBody
}

type GetHotspotPidListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotPidListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotPidListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotPidListResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotspotPidListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotspotPidListResponse) GetBody() *GetHotspotPidListResponseBody {
	return s.Body
}

func (s *GetHotspotPidListResponse) SetHeaders(v map[string]*string) *GetHotspotPidListResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotPidListResponse) SetStatusCode(v int32) *GetHotspotPidListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotPidListResponse) SetBody(v *GetHotspotPidListResponseBody) *GetHotspotPidListResponse {
	s.Body = v
	return s
}

func (s *GetHotspotPidListResponse) Validate() error {
	return dara.Validate(s)
}
