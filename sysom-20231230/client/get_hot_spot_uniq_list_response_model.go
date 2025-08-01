// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotSpotUniqListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotSpotUniqListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotSpotUniqListResponse
	GetStatusCode() *int32
	SetBody(v *GetHotSpotUniqListResponseBody) *GetHotSpotUniqListResponse
	GetBody() *GetHotSpotUniqListResponseBody
}

type GetHotSpotUniqListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotSpotUniqListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotSpotUniqListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotSpotUniqListResponse) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotSpotUniqListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotSpotUniqListResponse) GetBody() *GetHotSpotUniqListResponseBody {
	return s.Body
}

func (s *GetHotSpotUniqListResponse) SetHeaders(v map[string]*string) *GetHotSpotUniqListResponse {
	s.Headers = v
	return s
}

func (s *GetHotSpotUniqListResponse) SetStatusCode(v int32) *GetHotSpotUniqListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotSpotUniqListResponse) SetBody(v *GetHotSpotUniqListResponseBody) *GetHotSpotUniqListResponse {
	s.Body = v
	return s
}

func (s *GetHotSpotUniqListResponse) Validate() error {
	return dara.Validate(s)
}
