// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageLanguageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadMessageLanguageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadMessageLanguageResponse
	GetStatusCode() *int32
	SetBody(v *ReadMessageLanguageResponseBody) *ReadMessageLanguageResponse
	GetBody() *ReadMessageLanguageResponseBody
}

type ReadMessageLanguageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadMessageLanguageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadMessageLanguageResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageLanguageResponse) GoString() string {
	return s.String()
}

func (s *ReadMessageLanguageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadMessageLanguageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadMessageLanguageResponse) GetBody() *ReadMessageLanguageResponseBody {
	return s.Body
}

func (s *ReadMessageLanguageResponse) SetHeaders(v map[string]*string) *ReadMessageLanguageResponse {
	s.Headers = v
	return s
}

func (s *ReadMessageLanguageResponse) SetStatusCode(v int32) *ReadMessageLanguageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadMessageLanguageResponse) SetBody(v *ReadMessageLanguageResponseBody) *ReadMessageLanguageResponse {
	s.Body = v
	return s
}

func (s *ReadMessageLanguageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
