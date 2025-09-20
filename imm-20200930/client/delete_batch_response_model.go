// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBatchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBatchResponseBody) *DeleteBatchResponse
	GetBody() *DeleteBatchResponseBody
}

type DeleteBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBatchResponse) GoString() string {
	return s.String()
}

func (s *DeleteBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBatchResponse) GetBody() *DeleteBatchResponseBody {
	return s.Body
}

func (s *DeleteBatchResponse) SetHeaders(v map[string]*string) *DeleteBatchResponse {
	s.Headers = v
	return s
}

func (s *DeleteBatchResponse) SetStatusCode(v int32) *DeleteBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBatchResponse) SetBody(v *DeleteBatchResponseBody) *DeleteBatchResponse {
	s.Body = v
	return s
}

func (s *DeleteBatchResponse) Validate() error {
	return dara.Validate(s)
}
