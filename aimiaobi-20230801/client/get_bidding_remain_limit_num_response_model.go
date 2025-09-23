// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBiddingRemainLimitNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBiddingRemainLimitNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBiddingRemainLimitNumResponse
	GetStatusCode() *int32
	SetBody(v *GetBiddingRemainLimitNumResponseBody) *GetBiddingRemainLimitNumResponse
	GetBody() *GetBiddingRemainLimitNumResponseBody
}

type GetBiddingRemainLimitNumResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBiddingRemainLimitNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBiddingRemainLimitNumResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingRemainLimitNumResponse) GoString() string {
	return s.String()
}

func (s *GetBiddingRemainLimitNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBiddingRemainLimitNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBiddingRemainLimitNumResponse) GetBody() *GetBiddingRemainLimitNumResponseBody {
	return s.Body
}

func (s *GetBiddingRemainLimitNumResponse) SetHeaders(v map[string]*string) *GetBiddingRemainLimitNumResponse {
	s.Headers = v
	return s
}

func (s *GetBiddingRemainLimitNumResponse) SetStatusCode(v int32) *GetBiddingRemainLimitNumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBiddingRemainLimitNumResponse) SetBody(v *GetBiddingRemainLimitNumResponseBody) *GetBiddingRemainLimitNumResponse {
	s.Body = v
	return s
}

func (s *GetBiddingRemainLimitNumResponse) Validate() error {
	return dara.Validate(s)
}
