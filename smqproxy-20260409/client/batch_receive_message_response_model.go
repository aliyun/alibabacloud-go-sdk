// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchReceiveMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchReceiveMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchReceiveMessageResponse
	GetStatusCode() *int32
	SetBody(v *BatchReceiveMessageResponseBody) *BatchReceiveMessageResponse
	GetBody() *BatchReceiveMessageResponseBody
}

type BatchReceiveMessageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchReceiveMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchReceiveMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchReceiveMessageResponse) GoString() string {
	return s.String()
}

func (s *BatchReceiveMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchReceiveMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchReceiveMessageResponse) GetBody() *BatchReceiveMessageResponseBody {
	return s.Body
}

func (s *BatchReceiveMessageResponse) SetHeaders(v map[string]*string) *BatchReceiveMessageResponse {
	s.Headers = v
	return s
}

func (s *BatchReceiveMessageResponse) SetStatusCode(v int32) *BatchReceiveMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchReceiveMessageResponse) SetBody(v *BatchReceiveMessageResponseBody) *BatchReceiveMessageResponse {
	s.Body = v
	return s
}

func (s *BatchReceiveMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
