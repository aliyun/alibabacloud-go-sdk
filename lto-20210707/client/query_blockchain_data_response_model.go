// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlockchainDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBlockchainDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBlockchainDataResponse
	GetStatusCode() *int32
	SetBody(v *QueryBlockchainDataResponseBody) *QueryBlockchainDataResponse
	GetBody() *QueryBlockchainDataResponseBody
}

type QueryBlockchainDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBlockchainDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBlockchainDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainDataResponse) GoString() string {
	return s.String()
}

func (s *QueryBlockchainDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBlockchainDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBlockchainDataResponse) GetBody() *QueryBlockchainDataResponseBody {
	return s.Body
}

func (s *QueryBlockchainDataResponse) SetHeaders(v map[string]*string) *QueryBlockchainDataResponse {
	s.Headers = v
	return s
}

func (s *QueryBlockchainDataResponse) SetStatusCode(v int32) *QueryBlockchainDataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlockchainDataResponse) SetBody(v *QueryBlockchainDataResponseBody) *QueryBlockchainDataResponse {
	s.Body = v
	return s
}

func (s *QueryBlockchainDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
