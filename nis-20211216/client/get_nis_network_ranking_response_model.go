// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkRankingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNisNetworkRankingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNisNetworkRankingResponse
	GetStatusCode() *int32
	SetBody(v *GetNisNetworkRankingResponseBody) *GetNisNetworkRankingResponse
	GetBody() *GetNisNetworkRankingResponseBody
}

type GetNisNetworkRankingResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNisNetworkRankingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNisNetworkRankingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkRankingResponse) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNisNetworkRankingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNisNetworkRankingResponse) GetBody() *GetNisNetworkRankingResponseBody {
	return s.Body
}

func (s *GetNisNetworkRankingResponse) SetHeaders(v map[string]*string) *GetNisNetworkRankingResponse {
	s.Headers = v
	return s
}

func (s *GetNisNetworkRankingResponse) SetStatusCode(v int32) *GetNisNetworkRankingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNisNetworkRankingResponse) SetBody(v *GetNisNetworkRankingResponseBody) *GetNisNetworkRankingResponse {
	s.Body = v
	return s
}

func (s *GetNisNetworkRankingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
