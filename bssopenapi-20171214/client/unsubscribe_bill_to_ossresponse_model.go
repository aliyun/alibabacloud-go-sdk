// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeBillToOSSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnsubscribeBillToOSSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnsubscribeBillToOSSResponse
	GetStatusCode() *int32
	SetBody(v *UnsubscribeBillToOSSResponseBody) *UnsubscribeBillToOSSResponse
	GetBody() *UnsubscribeBillToOSSResponseBody
}

type UnsubscribeBillToOSSResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnsubscribeBillToOSSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnsubscribeBillToOSSResponse) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeBillToOSSResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeBillToOSSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnsubscribeBillToOSSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnsubscribeBillToOSSResponse) GetBody() *UnsubscribeBillToOSSResponseBody {
	return s.Body
}

func (s *UnsubscribeBillToOSSResponse) SetHeaders(v map[string]*string) *UnsubscribeBillToOSSResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeBillToOSSResponse) SetStatusCode(v int32) *UnsubscribeBillToOSSResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeBillToOSSResponse) SetBody(v *UnsubscribeBillToOSSResponseBody) *UnsubscribeBillToOSSResponse {
	s.Body = v
	return s
}

func (s *UnsubscribeBillToOSSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
