// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileProtectEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileProtectEventResponse
	GetStatusCode() *int32
	SetBody(v *ListFileProtectEventResponseBody) *ListFileProtectEventResponse
	GetBody() *ListFileProtectEventResponseBody
}

type ListFileProtectEventResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileProtectEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileProtectEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectEventResponse) GoString() string {
	return s.String()
}

func (s *ListFileProtectEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileProtectEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileProtectEventResponse) GetBody() *ListFileProtectEventResponseBody {
	return s.Body
}

func (s *ListFileProtectEventResponse) SetHeaders(v map[string]*string) *ListFileProtectEventResponse {
	s.Headers = v
	return s
}

func (s *ListFileProtectEventResponse) SetStatusCode(v int32) *ListFileProtectEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileProtectEventResponse) SetBody(v *ListFileProtectEventResponseBody) *ListFileProtectEventResponse {
	s.Body = v
	return s
}

func (s *ListFileProtectEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
