// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMqConsumerTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMqConsumerTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMqConsumerTagResponse
	GetStatusCode() *int32
	SetBody(v *GetMqConsumerTagResponseBody) *GetMqConsumerTagResponse
	GetBody() *GetMqConsumerTagResponseBody
}

type GetMqConsumerTagResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMqConsumerTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMqConsumerTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMqConsumerTagResponse) GoString() string {
	return s.String()
}

func (s *GetMqConsumerTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMqConsumerTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMqConsumerTagResponse) GetBody() *GetMqConsumerTagResponseBody {
	return s.Body
}

func (s *GetMqConsumerTagResponse) SetHeaders(v map[string]*string) *GetMqConsumerTagResponse {
	s.Headers = v
	return s
}

func (s *GetMqConsumerTagResponse) SetStatusCode(v int32) *GetMqConsumerTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqConsumerTagResponse) SetBody(v *GetMqConsumerTagResponseBody) *GetMqConsumerTagResponse {
	s.Body = v
	return s
}

func (s *GetMqConsumerTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
