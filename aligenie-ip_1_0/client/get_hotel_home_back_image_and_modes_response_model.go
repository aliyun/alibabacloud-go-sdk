// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelHomeBackImageAndModesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelHomeBackImageAndModesResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelHomeBackImageAndModesResponseBody) *GetHotelHomeBackImageAndModesResponse
	GetBody() *GetHotelHomeBackImageAndModesResponseBody
}

type GetHotelHomeBackImageAndModesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelHomeBackImageAndModesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelHomeBackImageAndModesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelHomeBackImageAndModesResponse) GoString() string {
	return s.String()
}

func (s *GetHotelHomeBackImageAndModesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelHomeBackImageAndModesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelHomeBackImageAndModesResponse) GetBody() *GetHotelHomeBackImageAndModesResponseBody {
	return s.Body
}

func (s *GetHotelHomeBackImageAndModesResponse) SetHeaders(v map[string]*string) *GetHotelHomeBackImageAndModesResponse {
	s.Headers = v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponse) SetStatusCode(v int32) *GetHotelHomeBackImageAndModesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponse) SetBody(v *GetHotelHomeBackImageAndModesResponseBody) *GetHotelHomeBackImageAndModesResponse {
	s.Body = v
	return s
}

func (s *GetHotelHomeBackImageAndModesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
