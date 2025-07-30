// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTensorboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTensorboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTensorboardResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTensorboardResponseBody) *DeleteTensorboardResponse
	GetBody() *DeleteTensorboardResponseBody
}

type DeleteTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTensorboardResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTensorboardResponse) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTensorboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTensorboardResponse) GetBody() *DeleteTensorboardResponseBody {
	return s.Body
}

func (s *DeleteTensorboardResponse) SetHeaders(v map[string]*string) *DeleteTensorboardResponse {
	s.Headers = v
	return s
}

func (s *DeleteTensorboardResponse) SetStatusCode(v int32) *DeleteTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTensorboardResponse) SetBody(v *DeleteTensorboardResponseBody) *DeleteTensorboardResponse {
	s.Body = v
	return s
}

func (s *DeleteTensorboardResponse) Validate() error {
	return dara.Validate(s)
}
