// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBlockchainMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBlockchainMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBlockchainMetadataResponse
	GetStatusCode() *int32
	SetBody(v *QueryBlockchainMetadataResponseBody) *QueryBlockchainMetadataResponse
	GetBody() *QueryBlockchainMetadataResponseBody
}

type QueryBlockchainMetadataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBlockchainMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBlockchainMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBlockchainMetadataResponse) GoString() string {
	return s.String()
}

func (s *QueryBlockchainMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBlockchainMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBlockchainMetadataResponse) GetBody() *QueryBlockchainMetadataResponseBody {
	return s.Body
}

func (s *QueryBlockchainMetadataResponse) SetHeaders(v map[string]*string) *QueryBlockchainMetadataResponse {
	s.Headers = v
	return s
}

func (s *QueryBlockchainMetadataResponse) SetStatusCode(v int32) *QueryBlockchainMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBlockchainMetadataResponse) SetBody(v *QueryBlockchainMetadataResponseBody) *QueryBlockchainMetadataResponse {
	s.Body = v
	return s
}

func (s *QueryBlockchainMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
