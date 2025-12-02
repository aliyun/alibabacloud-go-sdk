// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCallbackResponse
	GetStatusCode() *int32
	SetBody(v *QueryCallbackResponseBody) *QueryCallbackResponse
	GetBody() *QueryCallbackResponseBody
}

type QueryCallbackResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCallbackResponse) GoString() string {
	return s.String()
}

func (s *QueryCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCallbackResponse) GetBody() *QueryCallbackResponseBody {
	return s.Body
}

func (s *QueryCallbackResponse) SetHeaders(v map[string]*string) *QueryCallbackResponse {
	s.Headers = v
	return s
}

func (s *QueryCallbackResponse) SetStatusCode(v int32) *QueryCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallbackResponse) SetBody(v *QueryCallbackResponseBody) *QueryCallbackResponse {
	s.Body = v
	return s
}

func (s *QueryCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
