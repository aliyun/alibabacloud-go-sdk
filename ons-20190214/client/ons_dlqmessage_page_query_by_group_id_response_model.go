// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessagePageQueryByGroupIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsDLQMessagePageQueryByGroupIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsDLQMessagePageQueryByGroupIdResponse
	GetStatusCode() *int32
	SetBody(v *OnsDLQMessagePageQueryByGroupIdResponseBody) *OnsDLQMessagePageQueryByGroupIdResponse
	GetBody() *OnsDLQMessagePageQueryByGroupIdResponseBody
}

type OnsDLQMessagePageQueryByGroupIdResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsDLQMessagePageQueryByGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdResponse) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) GetBody() *OnsDLQMessagePageQueryByGroupIdResponseBody {
	return s.Body
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) SetHeaders(v map[string]*string) *OnsDLQMessagePageQueryByGroupIdResponse {
	s.Headers = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) SetStatusCode(v int32) *OnsDLQMessagePageQueryByGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) SetBody(v *OnsDLQMessagePageQueryByGroupIdResponseBody) *OnsDLQMessagePageQueryByGroupIdResponse {
	s.Body = v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
