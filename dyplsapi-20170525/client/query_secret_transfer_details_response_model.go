// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecretTransferDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySecretTransferDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySecretTransferDetailsResponse
	GetStatusCode() *int32
	SetBody(v *QuerySecretTransferDetailsResponseBody) *QuerySecretTransferDetailsResponse
	GetBody() *QuerySecretTransferDetailsResponseBody
}

type QuerySecretTransferDetailsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySecretTransferDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySecretTransferDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySecretTransferDetailsResponse) GoString() string {
	return s.String()
}

func (s *QuerySecretTransferDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySecretTransferDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySecretTransferDetailsResponse) GetBody() *QuerySecretTransferDetailsResponseBody {
	return s.Body
}

func (s *QuerySecretTransferDetailsResponse) SetHeaders(v map[string]*string) *QuerySecretTransferDetailsResponse {
	s.Headers = v
	return s
}

func (s *QuerySecretTransferDetailsResponse) SetStatusCode(v int32) *QuerySecretTransferDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySecretTransferDetailsResponse) SetBody(v *QuerySecretTransferDetailsResponseBody) *QuerySecretTransferDetailsResponse {
	s.Body = v
	return s
}

func (s *QuerySecretTransferDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
