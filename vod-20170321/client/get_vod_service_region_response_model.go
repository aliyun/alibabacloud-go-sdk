// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodServiceRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVodServiceRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVodServiceRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetVodServiceRegionResponseBody) *GetVodServiceRegionResponse
	GetBody() *GetVodServiceRegionResponseBody
}

type GetVodServiceRegionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVodServiceRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVodServiceRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVodServiceRegionResponse) GoString() string {
	return s.String()
}

func (s *GetVodServiceRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVodServiceRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVodServiceRegionResponse) GetBody() *GetVodServiceRegionResponseBody {
	return s.Body
}

func (s *GetVodServiceRegionResponse) SetHeaders(v map[string]*string) *GetVodServiceRegionResponse {
	s.Headers = v
	return s
}

func (s *GetVodServiceRegionResponse) SetStatusCode(v int32) *GetVodServiceRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVodServiceRegionResponse) SetBody(v *GetVodServiceRegionResponseBody) *GetVodServiceRegionResponse {
	s.Body = v
	return s
}

func (s *GetVodServiceRegionResponse) Validate() error {
	return dara.Validate(s)
}
