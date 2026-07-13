// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAtiRegistrantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAtiRegistrantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAtiRegistrantsResponse
	GetStatusCode() *int32
	SetBody(v *ListAtiRegistrantsResponseBody) *ListAtiRegistrantsResponse
	GetBody() *ListAtiRegistrantsResponseBody
}

type ListAtiRegistrantsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAtiRegistrantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAtiRegistrantsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAtiRegistrantsResponse) GoString() string {
	return s.String()
}

func (s *ListAtiRegistrantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAtiRegistrantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAtiRegistrantsResponse) GetBody() *ListAtiRegistrantsResponseBody {
	return s.Body
}

func (s *ListAtiRegistrantsResponse) SetHeaders(v map[string]*string) *ListAtiRegistrantsResponse {
	s.Headers = v
	return s
}

func (s *ListAtiRegistrantsResponse) SetStatusCode(v int32) *ListAtiRegistrantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAtiRegistrantsResponse) SetBody(v *ListAtiRegistrantsResponseBody) *ListAtiRegistrantsResponse {
	s.Body = v
	return s
}

func (s *ListAtiRegistrantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
