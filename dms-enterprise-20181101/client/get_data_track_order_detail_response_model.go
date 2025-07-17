// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataTrackOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataTrackOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDataTrackOrderDetailResponseBody) *GetDataTrackOrderDetailResponse
	GetBody() *GetDataTrackOrderDetailResponseBody
}

type GetDataTrackOrderDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataTrackOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataTrackOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDataTrackOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataTrackOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataTrackOrderDetailResponse) GetBody() *GetDataTrackOrderDetailResponseBody {
	return s.Body
}

func (s *GetDataTrackOrderDetailResponse) SetHeaders(v map[string]*string) *GetDataTrackOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDataTrackOrderDetailResponse) SetStatusCode(v int32) *GetDataTrackOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataTrackOrderDetailResponse) SetBody(v *GetDataTrackOrderDetailResponseBody) *GetDataTrackOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetDataTrackOrderDetailResponse) Validate() error {
	return dara.Validate(s)
}
