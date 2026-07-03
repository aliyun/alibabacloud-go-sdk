// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlertsResponse
	GetStatusCode() *int32
	SetBody(v *ListAlertsResponseBody) *ListAlertsResponse
	GetBody() *ListAlertsResponseBody
}

type ListAlertsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlertsResponse) GetBody() *ListAlertsResponseBody {
	return s.Body
}

func (s *ListAlertsResponse) SetHeaders(v map[string]*string) *ListAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertsResponse) SetStatusCode(v int32) *ListAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertsResponse) SetBody(v *ListAlertsResponseBody) *ListAlertsResponse {
	s.Body = v
	return s
}

func (s *ListAlertsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
