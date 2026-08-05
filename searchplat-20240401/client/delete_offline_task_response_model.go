// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfflineTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOfflineTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOfflineTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOfflineTaskResponseBody) *DeleteOfflineTaskResponse
	GetBody() *DeleteOfflineTaskResponseBody
}

type DeleteOfflineTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOfflineTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOfflineTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOfflineTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOfflineTaskResponse) GetBody() *DeleteOfflineTaskResponseBody {
	return s.Body
}

func (s *DeleteOfflineTaskResponse) SetHeaders(v map[string]*string) *DeleteOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOfflineTaskResponse) SetStatusCode(v int32) *DeleteOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOfflineTaskResponse) SetBody(v *DeleteOfflineTaskResponseBody) *DeleteOfflineTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteOfflineTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
