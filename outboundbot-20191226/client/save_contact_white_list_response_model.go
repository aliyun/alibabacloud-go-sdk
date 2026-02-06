// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveContactWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveContactWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *SaveContactWhiteListResponseBody) *SaveContactWhiteListResponse
	GetBody() *SaveContactWhiteListResponseBody
}

type SaveContactWhiteListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveContactWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveContactWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveContactWhiteListResponse) GoString() string {
	return s.String()
}

func (s *SaveContactWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveContactWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveContactWhiteListResponse) GetBody() *SaveContactWhiteListResponseBody {
	return s.Body
}

func (s *SaveContactWhiteListResponse) SetHeaders(v map[string]*string) *SaveContactWhiteListResponse {
	s.Headers = v
	return s
}

func (s *SaveContactWhiteListResponse) SetStatusCode(v int32) *SaveContactWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveContactWhiteListResponse) SetBody(v *SaveContactWhiteListResponseBody) *SaveContactWhiteListResponse {
	s.Body = v
	return s
}

func (s *SaveContactWhiteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
