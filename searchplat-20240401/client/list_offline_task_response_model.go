// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfflineTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOfflineTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOfflineTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListOfflineTaskResponseBody) *ListOfflineTaskResponse
	GetBody() *ListOfflineTaskResponseBody
}

type ListOfflineTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOfflineTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOfflineTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOfflineTaskResponse) GetBody() *ListOfflineTaskResponseBody {
	return s.Body
}

func (s *ListOfflineTaskResponse) SetHeaders(v map[string]*string) *ListOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOfflineTaskResponse) SetStatusCode(v int32) *ListOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOfflineTaskResponse) SetBody(v *ListOfflineTaskResponseBody) *ListOfflineTaskResponse {
	s.Body = v
	return s
}

func (s *ListOfflineTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
