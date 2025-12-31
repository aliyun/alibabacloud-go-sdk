// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveContactsResponse
	GetStatusCode() *int32
	SetBody(v *SaveContactsResponseBody) *SaveContactsResponse
	GetBody() *SaveContactsResponseBody
}

type SaveContactsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveContactsResponse) GoString() string {
	return s.String()
}

func (s *SaveContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveContactsResponse) GetBody() *SaveContactsResponseBody {
	return s.Body
}

func (s *SaveContactsResponse) SetHeaders(v map[string]*string) *SaveContactsResponse {
	s.Headers = v
	return s
}

func (s *SaveContactsResponse) SetStatusCode(v int32) *SaveContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveContactsResponse) SetBody(v *SaveContactsResponseBody) *SaveContactsResponse {
	s.Body = v
	return s
}

func (s *SaveContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
