// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMessageByMessageIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMessageByMessageIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMessageByMessageIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryMessageByMessageIdResponseBody) *QueryMessageByMessageIdResponse
	GetBody() *QueryMessageByMessageIdResponseBody
}

type QueryMessageByMessageIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMessageByMessageIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMessageByMessageIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMessageByMessageIdResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageByMessageIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMessageByMessageIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMessageByMessageIdResponse) GetBody() *QueryMessageByMessageIdResponseBody {
	return s.Body
}

func (s *QueryMessageByMessageIdResponse) SetHeaders(v map[string]*string) *QueryMessageByMessageIdResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageByMessageIdResponse) SetStatusCode(v int32) *QueryMessageByMessageIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageByMessageIdResponse) SetBody(v *QueryMessageByMessageIdResponseBody) *QueryMessageByMessageIdResponse {
	s.Body = v
	return s
}

func (s *QueryMessageByMessageIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
