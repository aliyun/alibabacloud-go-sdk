// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastVideosByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBroadcastVideosByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBroadcastVideosByIdResponse
	GetStatusCode() *int32
	SetBody(v *ListBroadcastVideosByIdResponseBody) *ListBroadcastVideosByIdResponse
	GetBody() *ListBroadcastVideosByIdResponseBody
}

type ListBroadcastVideosByIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBroadcastVideosByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBroadcastVideosByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastVideosByIdResponse) GoString() string {
	return s.String()
}

func (s *ListBroadcastVideosByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBroadcastVideosByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBroadcastVideosByIdResponse) GetBody() *ListBroadcastVideosByIdResponseBody {
	return s.Body
}

func (s *ListBroadcastVideosByIdResponse) SetHeaders(v map[string]*string) *ListBroadcastVideosByIdResponse {
	s.Headers = v
	return s
}

func (s *ListBroadcastVideosByIdResponse) SetStatusCode(v int32) *ListBroadcastVideosByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBroadcastVideosByIdResponse) SetBody(v *ListBroadcastVideosByIdResponseBody) *ListBroadcastVideosByIdResponse {
	s.Body = v
	return s
}

func (s *ListBroadcastVideosByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
