// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReceiverDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReceiverDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReceiverDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryReceiverDetailResponseBody) *QueryReceiverDetailResponse
	GetBody() *QueryReceiverDetailResponseBody
}

type QueryReceiverDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReceiverDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReceiverDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryReceiverDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReceiverDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReceiverDetailResponse) GetBody() *QueryReceiverDetailResponseBody {
	return s.Body
}

func (s *QueryReceiverDetailResponse) SetHeaders(v map[string]*string) *QueryReceiverDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryReceiverDetailResponse) SetStatusCode(v int32) *QueryReceiverDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReceiverDetailResponse) SetBody(v *QueryReceiverDetailResponseBody) *QueryReceiverDetailResponse {
	s.Body = v
	return s
}

func (s *QueryReceiverDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
