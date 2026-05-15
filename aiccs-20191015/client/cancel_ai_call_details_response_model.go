// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAiCallDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAiCallDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAiCallDetailsResponse
	GetStatusCode() *int32
	SetBody(v *CancelAiCallDetailsResponseBody) *CancelAiCallDetailsResponse
	GetBody() *CancelAiCallDetailsResponseBody
}

type CancelAiCallDetailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAiCallDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAiCallDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAiCallDetailsResponse) GoString() string {
	return s.String()
}

func (s *CancelAiCallDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAiCallDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAiCallDetailsResponse) GetBody() *CancelAiCallDetailsResponseBody {
	return s.Body
}

func (s *CancelAiCallDetailsResponse) SetHeaders(v map[string]*string) *CancelAiCallDetailsResponse {
	s.Headers = v
	return s
}

func (s *CancelAiCallDetailsResponse) SetStatusCode(v int32) *CancelAiCallDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAiCallDetailsResponse) SetBody(v *CancelAiCallDetailsResponseBody) *CancelAiCallDetailsResponse {
	s.Body = v
	return s
}

func (s *CancelAiCallDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
