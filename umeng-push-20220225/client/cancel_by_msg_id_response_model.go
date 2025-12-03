// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelByMsgIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelByMsgIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelByMsgIdResponse
	GetStatusCode() *int32
	SetBody(v *CancelByMsgIdResponseBody) *CancelByMsgIdResponse
	GetBody() *CancelByMsgIdResponseBody
}

type CancelByMsgIdResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelByMsgIdResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *CancelByMsgIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelByMsgIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelByMsgIdResponse) GetBody() *CancelByMsgIdResponseBody {
	return s.Body
}

func (s *CancelByMsgIdResponse) SetHeaders(v map[string]*string) *CancelByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *CancelByMsgIdResponse) SetStatusCode(v int32) *CancelByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelByMsgIdResponse) SetBody(v *CancelByMsgIdResponseBody) *CancelByMsgIdResponse {
	s.Body = v
	return s
}

func (s *CancelByMsgIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
