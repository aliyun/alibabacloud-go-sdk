// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineModelResponse
	GetStatusCode() *int32
	SetBody(v *OnlineModelResponseBody) *OnlineModelResponse
	GetBody() *OnlineModelResponseBody
}

type OnlineModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineModelResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineModelResponse) GoString() string {
	return s.String()
}

func (s *OnlineModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineModelResponse) GetBody() *OnlineModelResponseBody {
	return s.Body
}

func (s *OnlineModelResponse) SetHeaders(v map[string]*string) *OnlineModelResponse {
	s.Headers = v
	return s
}

func (s *OnlineModelResponse) SetStatusCode(v int32) *OnlineModelResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineModelResponse) SetBody(v *OnlineModelResponseBody) *OnlineModelResponse {
	s.Body = v
	return s
}

func (s *OnlineModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
