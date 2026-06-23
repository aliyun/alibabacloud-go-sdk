// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSendMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSendMessageResponse
	GetStatusCode() *int32
	SetBody(v *BatchSendMessageResponseBody) *BatchSendMessageResponse
	GetBody() *BatchSendMessageResponseBody
}

type BatchSendMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSendMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageResponse) GoString() string {
	return s.String()
}

func (s *BatchSendMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSendMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSendMessageResponse) GetBody() *BatchSendMessageResponseBody {
	return s.Body
}

func (s *BatchSendMessageResponse) SetHeaders(v map[string]*string) *BatchSendMessageResponse {
	s.Headers = v
	return s
}

func (s *BatchSendMessageResponse) SetStatusCode(v int32) *BatchSendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendMessageResponse) SetBody(v *BatchSendMessageResponseBody) *BatchSendMessageResponse {
	s.Body = v
	return s
}

func (s *BatchSendMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
