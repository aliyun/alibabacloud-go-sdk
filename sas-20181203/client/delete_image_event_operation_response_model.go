// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageEventOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImageEventOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImageEventOperationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImageEventOperationResponseBody) *DeleteImageEventOperationResponse
	GetBody() *DeleteImageEventOperationResponseBody
}

type DeleteImageEventOperationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageEventOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageEventOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageEventOperationResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageEventOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImageEventOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImageEventOperationResponse) GetBody() *DeleteImageEventOperationResponseBody {
	return s.Body
}

func (s *DeleteImageEventOperationResponse) SetHeaders(v map[string]*string) *DeleteImageEventOperationResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageEventOperationResponse) SetStatusCode(v int32) *DeleteImageEventOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageEventOperationResponse) SetBody(v *DeleteImageEventOperationResponseBody) *DeleteImageEventOperationResponse {
	s.Body = v
	return s
}

func (s *DeleteImageEventOperationResponse) Validate() error {
	return dara.Validate(s)
}
