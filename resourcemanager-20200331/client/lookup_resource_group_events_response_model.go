// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLookupResourceGroupEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LookupResourceGroupEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LookupResourceGroupEventsResponse
	GetStatusCode() *int32
	SetBody(v *LookupResourceGroupEventsResponseBody) *LookupResourceGroupEventsResponse
	GetBody() *LookupResourceGroupEventsResponseBody
}

type LookupResourceGroupEventsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LookupResourceGroupEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LookupResourceGroupEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s LookupResourceGroupEventsResponse) GoString() string {
	return s.String()
}

func (s *LookupResourceGroupEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LookupResourceGroupEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LookupResourceGroupEventsResponse) GetBody() *LookupResourceGroupEventsResponseBody {
	return s.Body
}

func (s *LookupResourceGroupEventsResponse) SetHeaders(v map[string]*string) *LookupResourceGroupEventsResponse {
	s.Headers = v
	return s
}

func (s *LookupResourceGroupEventsResponse) SetStatusCode(v int32) *LookupResourceGroupEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupResourceGroupEventsResponse) SetBody(v *LookupResourceGroupEventsResponseBody) *LookupResourceGroupEventsResponse {
	s.Body = v
	return s
}

func (s *LookupResourceGroupEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
