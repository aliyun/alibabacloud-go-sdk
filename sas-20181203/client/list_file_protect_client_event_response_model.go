// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileProtectClientEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileProtectClientEventResponse
	GetStatusCode() *int32
	SetBody(v *ListFileProtectClientEventResponseBody) *ListFileProtectClientEventResponse
	GetBody() *ListFileProtectClientEventResponseBody
}

type ListFileProtectClientEventResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileProtectClientEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileProtectClientEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientEventResponse) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileProtectClientEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileProtectClientEventResponse) GetBody() *ListFileProtectClientEventResponseBody {
	return s.Body
}

func (s *ListFileProtectClientEventResponse) SetHeaders(v map[string]*string) *ListFileProtectClientEventResponse {
	s.Headers = v
	return s
}

func (s *ListFileProtectClientEventResponse) SetStatusCode(v int32) *ListFileProtectClientEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileProtectClientEventResponse) SetBody(v *ListFileProtectClientEventResponseBody) *ListFileProtectClientEventResponse {
	s.Body = v
	return s
}

func (s *ListFileProtectClientEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
