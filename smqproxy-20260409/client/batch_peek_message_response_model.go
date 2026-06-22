// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPeekMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchPeekMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchPeekMessageResponse
	GetStatusCode() *int32
	SetBody(v *BatchPeekMessageResponseBody) *BatchPeekMessageResponse
	GetBody() *BatchPeekMessageResponseBody
}

type BatchPeekMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchPeekMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchPeekMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchPeekMessageResponse) GoString() string {
	return s.String()
}

func (s *BatchPeekMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchPeekMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchPeekMessageResponse) GetBody() *BatchPeekMessageResponseBody {
	return s.Body
}

func (s *BatchPeekMessageResponse) SetHeaders(v map[string]*string) *BatchPeekMessageResponse {
	s.Headers = v
	return s
}

func (s *BatchPeekMessageResponse) SetStatusCode(v int32) *BatchPeekMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchPeekMessageResponse) SetBody(v *BatchPeekMessageResponseBody) *BatchPeekMessageResponse {
	s.Body = v
	return s
}

func (s *BatchPeekMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
