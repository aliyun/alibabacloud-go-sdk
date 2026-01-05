// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSlotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSlotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSlotsResponse
	GetStatusCode() *int32
	SetBody(v *ListSlotsResponseBody) *ListSlotsResponse
	GetBody() *ListSlotsResponseBody
}

type ListSlotsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSlotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSlotsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSlotsResponse) GoString() string {
	return s.String()
}

func (s *ListSlotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSlotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSlotsResponse) GetBody() *ListSlotsResponseBody {
	return s.Body
}

func (s *ListSlotsResponse) SetHeaders(v map[string]*string) *ListSlotsResponse {
	s.Headers = v
	return s
}

func (s *ListSlotsResponse) SetStatusCode(v int32) *ListSlotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlotsResponse) SetBody(v *ListSlotsResponseBody) *ListSlotsResponse {
	s.Body = v
	return s
}

func (s *ListSlotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
