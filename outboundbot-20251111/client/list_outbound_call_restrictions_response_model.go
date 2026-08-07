// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundCallRestrictionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOutboundCallRestrictionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOutboundCallRestrictionsResponse
	GetStatusCode() *int32
	SetBody(v *ListOutboundCallRestrictionsResponseBody) *ListOutboundCallRestrictionsResponse
	GetBody() *ListOutboundCallRestrictionsResponseBody
}

type ListOutboundCallRestrictionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOutboundCallRestrictionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOutboundCallRestrictionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundCallRestrictionsResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundCallRestrictionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOutboundCallRestrictionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOutboundCallRestrictionsResponse) GetBody() *ListOutboundCallRestrictionsResponseBody {
	return s.Body
}

func (s *ListOutboundCallRestrictionsResponse) SetHeaders(v map[string]*string) *ListOutboundCallRestrictionsResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundCallRestrictionsResponse) SetStatusCode(v int32) *ListOutboundCallRestrictionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOutboundCallRestrictionsResponse) SetBody(v *ListOutboundCallRestrictionsResponseBody) *ListOutboundCallRestrictionsResponse {
	s.Body = v
	return s
}

func (s *ListOutboundCallRestrictionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
