// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryYuqingMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryYuqingMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryYuqingMessageResponse
	GetStatusCode() *int32
	SetBody(v *QueryYuqingMessageResponseBody) *QueryYuqingMessageResponse
	GetBody() *QueryYuqingMessageResponseBody
}

type QueryYuqingMessageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryYuqingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryYuqingMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryYuqingMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryYuqingMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryYuqingMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryYuqingMessageResponse) GetBody() *QueryYuqingMessageResponseBody {
	return s.Body
}

func (s *QueryYuqingMessageResponse) SetHeaders(v map[string]*string) *QueryYuqingMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryYuqingMessageResponse) SetStatusCode(v int32) *QueryYuqingMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryYuqingMessageResponse) SetBody(v *QueryYuqingMessageResponseBody) *QueryYuqingMessageResponse {
	s.Body = v
	return s
}

func (s *QueryYuqingMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
