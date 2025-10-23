// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchSendMailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchSendMailResponse
	GetStatusCode() *int32
	SetBody(v *BatchSendMailResponseBody) *BatchSendMailResponse
	GetBody() *BatchSendMailResponseBody
}

type BatchSendMailResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSendMailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSendMailResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMailResponse) GoString() string {
	return s.String()
}

func (s *BatchSendMailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchSendMailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchSendMailResponse) GetBody() *BatchSendMailResponseBody {
	return s.Body
}

func (s *BatchSendMailResponse) SetHeaders(v map[string]*string) *BatchSendMailResponse {
	s.Headers = v
	return s
}

func (s *BatchSendMailResponse) SetStatusCode(v int32) *BatchSendMailResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSendMailResponse) SetBody(v *BatchSendMailResponseBody) *BatchSendMailResponse {
	s.Body = v
	return s
}

func (s *BatchSendMailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
