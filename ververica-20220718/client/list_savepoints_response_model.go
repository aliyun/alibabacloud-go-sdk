// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSavepointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSavepointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSavepointsResponse
	GetStatusCode() *int32
	SetBody(v *ListSavepointsResponseBody) *ListSavepointsResponse
	GetBody() *ListSavepointsResponseBody
}

type ListSavepointsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSavepointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSavepointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSavepointsResponse) GoString() string {
	return s.String()
}

func (s *ListSavepointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSavepointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSavepointsResponse) GetBody() *ListSavepointsResponseBody {
	return s.Body
}

func (s *ListSavepointsResponse) SetHeaders(v map[string]*string) *ListSavepointsResponse {
	s.Headers = v
	return s
}

func (s *ListSavepointsResponse) SetStatusCode(v int32) *ListSavepointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSavepointsResponse) SetBody(v *ListSavepointsResponseBody) *ListSavepointsResponse {
	s.Body = v
	return s
}

func (s *ListSavepointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
