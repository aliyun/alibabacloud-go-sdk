// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdnSubTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCdnSubTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCdnSubTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCdnSubTaskResponseBody) *DeleteCdnSubTaskResponse
	GetBody() *DeleteCdnSubTaskResponseBody
}

type DeleteCdnSubTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCdnSubTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCdnSubTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdnSubTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteCdnSubTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCdnSubTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCdnSubTaskResponse) GetBody() *DeleteCdnSubTaskResponseBody {
	return s.Body
}

func (s *DeleteCdnSubTaskResponse) SetHeaders(v map[string]*string) *DeleteCdnSubTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteCdnSubTaskResponse) SetStatusCode(v int32) *DeleteCdnSubTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCdnSubTaskResponse) SetBody(v *DeleteCdnSubTaskResponseBody) *DeleteCdnSubTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteCdnSubTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
