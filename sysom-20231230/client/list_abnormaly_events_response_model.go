// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAbnormalyEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAbnormalyEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAbnormalyEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListAbnormalyEventsResponseBody) *ListAbnormalyEventsResponse
	GetBody() *ListAbnormalyEventsResponseBody
}

type ListAbnormalyEventsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAbnormalyEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAbnormalyEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAbnormalyEventsResponse) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAbnormalyEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAbnormalyEventsResponse) GetBody() *ListAbnormalyEventsResponseBody {
	return s.Body
}

func (s *ListAbnormalyEventsResponse) SetHeaders(v map[string]*string) *ListAbnormalyEventsResponse {
	s.Headers = v
	return s
}

func (s *ListAbnormalyEventsResponse) SetStatusCode(v int32) *ListAbnormalyEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAbnormalyEventsResponse) SetBody(v *ListAbnormalyEventsResponseBody) *ListAbnormalyEventsResponse {
	s.Body = v
	return s
}

func (s *ListAbnormalyEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
