// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionZoneProxyPatternResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRecursionZoneProxyPatternResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRecursionZoneProxyPatternResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRecursionZoneProxyPatternResponseBody) *UpdateRecursionZoneProxyPatternResponse
	GetBody() *UpdateRecursionZoneProxyPatternResponseBody
}

type UpdateRecursionZoneProxyPatternResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecursionZoneProxyPatternResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRecursionZoneProxyPatternResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionZoneProxyPatternResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecursionZoneProxyPatternResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRecursionZoneProxyPatternResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRecursionZoneProxyPatternResponse) GetBody() *UpdateRecursionZoneProxyPatternResponseBody {
	return s.Body
}

func (s *UpdateRecursionZoneProxyPatternResponse) SetHeaders(v map[string]*string) *UpdateRecursionZoneProxyPatternResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecursionZoneProxyPatternResponse) SetStatusCode(v int32) *UpdateRecursionZoneProxyPatternResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecursionZoneProxyPatternResponse) SetBody(v *UpdateRecursionZoneProxyPatternResponseBody) *UpdateRecursionZoneProxyPatternResponse {
	s.Body = v
	return s
}

func (s *UpdateRecursionZoneProxyPatternResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
