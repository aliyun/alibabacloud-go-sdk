// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActivatedAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListActivatedAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListActivatedAlertsResponse
	GetStatusCode() *int32
	SetBody(v *ListActivatedAlertsResponseBody) *ListActivatedAlertsResponse
	GetBody() *ListActivatedAlertsResponseBody
}

type ListActivatedAlertsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActivatedAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActivatedAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListActivatedAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListActivatedAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListActivatedAlertsResponse) GetBody() *ListActivatedAlertsResponseBody {
	return s.Body
}

func (s *ListActivatedAlertsResponse) SetHeaders(v map[string]*string) *ListActivatedAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListActivatedAlertsResponse) SetStatusCode(v int32) *ListActivatedAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActivatedAlertsResponse) SetBody(v *ListActivatedAlertsResponseBody) *ListActivatedAlertsResponse {
	s.Body = v
	return s
}

func (s *ListActivatedAlertsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
