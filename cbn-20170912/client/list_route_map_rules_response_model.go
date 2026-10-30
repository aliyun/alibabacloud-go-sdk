// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRouteMapRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRouteMapRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRouteMapRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListRouteMapRulesResponseBody) *ListRouteMapRulesResponse
	GetBody() *ListRouteMapRulesResponseBody
}

type ListRouteMapRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRouteMapRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRouteMapRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRouteMapRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRouteMapRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRouteMapRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRouteMapRulesResponse) GetBody() *ListRouteMapRulesResponseBody {
	return s.Body
}

func (s *ListRouteMapRulesResponse) SetHeaders(v map[string]*string) *ListRouteMapRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRouteMapRulesResponse) SetStatusCode(v int32) *ListRouteMapRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRouteMapRulesResponse) SetBody(v *ListRouteMapRulesResponseBody) *ListRouteMapRulesResponse {
	s.Body = v
	return s
}

func (s *ListRouteMapRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
