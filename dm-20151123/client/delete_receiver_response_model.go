// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReceiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteReceiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteReceiverResponse
	GetStatusCode() *int32
	SetBody(v *DeleteReceiverResponseBody) *DeleteReceiverResponse
	GetBody() *DeleteReceiverResponseBody
}

type DeleteReceiverResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteReceiverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteReceiverResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteReceiverResponse) GoString() string {
	return s.String()
}

func (s *DeleteReceiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteReceiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteReceiverResponse) GetBody() *DeleteReceiverResponseBody {
	return s.Body
}

func (s *DeleteReceiverResponse) SetHeaders(v map[string]*string) *DeleteReceiverResponse {
	s.Headers = v
	return s
}

func (s *DeleteReceiverResponse) SetStatusCode(v int32) *DeleteReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteReceiverResponse) SetBody(v *DeleteReceiverResponseBody) *DeleteReceiverResponse {
	s.Body = v
	return s
}

func (s *DeleteReceiverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
