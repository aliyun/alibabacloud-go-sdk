// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribePbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubscribePbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubscribePbcResponse
	GetStatusCode() *int32
	SetBody(v *SubscribePbcResponseBody) *SubscribePbcResponse
	GetBody() *SubscribePbcResponseBody
}

type SubscribePbcResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribePbcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribePbcResponse) String() string {
	return dara.Prettify(s)
}

func (s SubscribePbcResponse) GoString() string {
	return s.String()
}

func (s *SubscribePbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubscribePbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubscribePbcResponse) GetBody() *SubscribePbcResponseBody {
	return s.Body
}

func (s *SubscribePbcResponse) SetHeaders(v map[string]*string) *SubscribePbcResponse {
	s.Headers = v
	return s
}

func (s *SubscribePbcResponse) SetStatusCode(v int32) *SubscribePbcResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribePbcResponse) SetBody(v *SubscribePbcResponseBody) *SubscribePbcResponse {
	s.Body = v
	return s
}

func (s *SubscribePbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
