// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSampleUtterancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelSampleUtterancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelSampleUtterancesResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelSampleUtterancesResponseBody) *GetHotelSampleUtterancesResponse
	GetBody() *GetHotelSampleUtterancesResponseBody
}

type GetHotelSampleUtterancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelSampleUtterancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelSampleUtterancesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSampleUtterancesResponse) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelSampleUtterancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelSampleUtterancesResponse) GetBody() *GetHotelSampleUtterancesResponseBody {
	return s.Body
}

func (s *GetHotelSampleUtterancesResponse) SetHeaders(v map[string]*string) *GetHotelSampleUtterancesResponse {
	s.Headers = v
	return s
}

func (s *GetHotelSampleUtterancesResponse) SetStatusCode(v int32) *GetHotelSampleUtterancesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSampleUtterancesResponse) SetBody(v *GetHotelSampleUtterancesResponseBody) *GetHotelSampleUtterancesResponse {
	s.Body = v
	return s
}

func (s *GetHotelSampleUtterancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
