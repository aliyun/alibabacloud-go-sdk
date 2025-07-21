// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReceiverDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReceiverDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReceiverDetailResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReceiverDetailResponseBody) *DeleteReceiverDetailResponse
	GetBody() *DeleteReceiverDetailResponseBody
}

type DeleteReceiverDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReceiverDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReceiverDetailResponse) GoString() string {
	return s.String()
}

func (s *DeleteReceiverDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReceiverDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReceiverDetailResponse) GetBody() *DeleteReceiverDetailResponseBody {
	return s.Body
}

func (s *DeleteReceiverDetailResponse) SetHeaders(v map[string]*string) *DeleteReceiverDetailResponse {
	s.Headers = v
	return s
}

func (s *DeleteReceiverDetailResponse) SetStatusCode(v int32) *DeleteReceiverDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReceiverDetailResponse) SetBody(v *DeleteReceiverDetailResponseBody) *DeleteReceiverDetailResponse {
	s.Body = v
	return s
}

func (s *DeleteReceiverDetailResponse) Validate() error {
	return dara.Validate(s)
}
