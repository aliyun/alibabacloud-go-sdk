// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeBillToOSSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubscribeBillToOSSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubscribeBillToOSSResponse
	GetStatusCode() *int32
	SetBody(v *SubscribeBillToOSSResponseBody) *SubscribeBillToOSSResponse
	GetBody() *SubscribeBillToOSSResponseBody
}

type SubscribeBillToOSSResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeBillToOSSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeBillToOSSResponse) String() string {
	return dara.Prettify(s)
}

func (s SubscribeBillToOSSResponse) GoString() string {
	return s.String()
}

func (s *SubscribeBillToOSSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubscribeBillToOSSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubscribeBillToOSSResponse) GetBody() *SubscribeBillToOSSResponseBody {
	return s.Body
}

func (s *SubscribeBillToOSSResponse) SetHeaders(v map[string]*string) *SubscribeBillToOSSResponse {
	s.Headers = v
	return s
}

func (s *SubscribeBillToOSSResponse) SetStatusCode(v int32) *SubscribeBillToOSSResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeBillToOSSResponse) SetBody(v *SubscribeBillToOSSResponseBody) *SubscribeBillToOSSResponse {
	s.Body = v
	return s
}

func (s *SubscribeBillToOSSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
