// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMgsApirestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveMgsApirestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveMgsApirestResponse
	GetStatusCode() *int32
	SetBody(v *SaveMgsApirestResponseBody) *SaveMgsApirestResponse
	GetBody() *SaveMgsApirestResponseBody
}

type SaveMgsApirestResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveMgsApirestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveMgsApirestResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveMgsApirestResponse) GoString() string {
	return s.String()
}

func (s *SaveMgsApirestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveMgsApirestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveMgsApirestResponse) GetBody() *SaveMgsApirestResponseBody {
	return s.Body
}

func (s *SaveMgsApirestResponse) SetHeaders(v map[string]*string) *SaveMgsApirestResponse {
	s.Headers = v
	return s
}

func (s *SaveMgsApirestResponse) SetStatusCode(v int32) *SaveMgsApirestResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveMgsApirestResponse) SetBody(v *SaveMgsApirestResponseBody) *SaveMgsApirestResponse {
	s.Body = v
	return s
}

func (s *SaveMgsApirestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
