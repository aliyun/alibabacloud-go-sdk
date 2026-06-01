// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayHistoryServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRayHistoryServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRayHistoryServersResponse
	GetStatusCode() *int32
	SetBody(v *ListRayHistoryServersResponseBody) *ListRayHistoryServersResponse
	GetBody() *ListRayHistoryServersResponseBody
}

type ListRayHistoryServersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRayHistoryServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRayHistoryServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRayHistoryServersResponse) GoString() string {
	return s.String()
}

func (s *ListRayHistoryServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRayHistoryServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRayHistoryServersResponse) GetBody() *ListRayHistoryServersResponseBody {
	return s.Body
}

func (s *ListRayHistoryServersResponse) SetHeaders(v map[string]*string) *ListRayHistoryServersResponse {
	s.Headers = v
	return s
}

func (s *ListRayHistoryServersResponse) SetStatusCode(v int32) *ListRayHistoryServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRayHistoryServersResponse) SetBody(v *ListRayHistoryServersResponseBody) *ListRayHistoryServersResponse {
	s.Body = v
	return s
}

func (s *ListRayHistoryServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
