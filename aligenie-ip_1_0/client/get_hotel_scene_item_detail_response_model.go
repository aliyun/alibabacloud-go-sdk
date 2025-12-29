// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSceneItemDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelSceneItemDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelSceneItemDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelSceneItemDetailResponseBody) *GetHotelSceneItemDetailResponse
	GetBody() *GetHotelSceneItemDetailResponseBody
}

type GetHotelSceneItemDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelSceneItemDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelSceneItemDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSceneItemDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotelSceneItemDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelSceneItemDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelSceneItemDetailResponse) GetBody() *GetHotelSceneItemDetailResponseBody {
	return s.Body
}

func (s *GetHotelSceneItemDetailResponse) SetHeaders(v map[string]*string) *GetHotelSceneItemDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotelSceneItemDetailResponse) SetStatusCode(v int32) *GetHotelSceneItemDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSceneItemDetailResponse) SetBody(v *GetHotelSceneItemDetailResponseBody) *GetHotelSceneItemDetailResponse {
	s.Body = v
	return s
}

func (s *GetHotelSceneItemDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
