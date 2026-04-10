// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *QueryApiKeysResponseBody) *QueryApiKeysResponse
	GetBody() *QueryApiKeysResponseBody
}

type QueryApiKeysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryApiKeysResponse) GoString() string {
	return s.String()
}

func (s *QueryApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryApiKeysResponse) GetBody() *QueryApiKeysResponseBody {
	return s.Body
}

func (s *QueryApiKeysResponse) SetHeaders(v map[string]*string) *QueryApiKeysResponse {
	s.Headers = v
	return s
}

func (s *QueryApiKeysResponse) SetStatusCode(v int32) *QueryApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApiKeysResponse) SetBody(v *QueryApiKeysResponseBody) *QueryApiKeysResponse {
	s.Body = v
	return s
}

func (s *QueryApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
