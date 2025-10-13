// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpotPriceHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSpotPriceHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSpotPriceHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetSpotPriceHistoryResponseBody) *GetSpotPriceHistoryResponse
	GetBody() *GetSpotPriceHistoryResponseBody
}

type GetSpotPriceHistoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpotPriceHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpotPriceHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSpotPriceHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetSpotPriceHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSpotPriceHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSpotPriceHistoryResponse) GetBody() *GetSpotPriceHistoryResponseBody {
	return s.Body
}

func (s *GetSpotPriceHistoryResponse) SetHeaders(v map[string]*string) *GetSpotPriceHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetSpotPriceHistoryResponse) SetStatusCode(v int32) *GetSpotPriceHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpotPriceHistoryResponse) SetBody(v *GetSpotPriceHistoryResponseBody) *GetSpotPriceHistoryResponse {
	s.Body = v
	return s
}

func (s *GetSpotPriceHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
