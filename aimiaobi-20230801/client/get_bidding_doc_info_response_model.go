// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBiddingDocInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBiddingDocInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBiddingDocInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBiddingDocInfoResponseBody) *GetBiddingDocInfoResponse
	GetBody() *GetBiddingDocInfoResponseBody
}

type GetBiddingDocInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBiddingDocInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBiddingDocInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBiddingDocInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBiddingDocInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBiddingDocInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBiddingDocInfoResponse) GetBody() *GetBiddingDocInfoResponseBody {
	return s.Body
}

func (s *GetBiddingDocInfoResponse) SetHeaders(v map[string]*string) *GetBiddingDocInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBiddingDocInfoResponse) SetStatusCode(v int32) *GetBiddingDocInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBiddingDocInfoResponse) SetBody(v *GetBiddingDocInfoResponseBody) *GetBiddingDocInfoResponse {
	s.Body = v
	return s
}

func (s *GetBiddingDocInfoResponse) Validate() error {
	return dara.Validate(s)
}
