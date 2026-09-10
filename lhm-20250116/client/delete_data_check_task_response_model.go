// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataCheckTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataCheckTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataCheckTaskResponseBody) *DeleteDataCheckTaskResponse
	GetBody() *DeleteDataCheckTaskResponseBody
}

type DeleteDataCheckTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataCheckTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataCheckTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataCheckTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataCheckTaskResponse) GetBody() *DeleteDataCheckTaskResponseBody {
	return s.Body
}

func (s *DeleteDataCheckTaskResponse) SetHeaders(v map[string]*string) *DeleteDataCheckTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataCheckTaskResponse) SetStatusCode(v int32) *DeleteDataCheckTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataCheckTaskResponse) SetBody(v *DeleteDataCheckTaskResponseBody) *DeleteDataCheckTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteDataCheckTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
