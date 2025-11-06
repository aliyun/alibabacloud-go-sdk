// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAckInfoOfMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAckInfoOfMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAckInfoOfMessageResponse
	GetStatusCode() *int32
	SetBody(v *GetAckInfoOfMessageResponseBody) *GetAckInfoOfMessageResponse
	GetBody() *GetAckInfoOfMessageResponseBody
}

type GetAckInfoOfMessageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAckInfoOfMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAckInfoOfMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAckInfoOfMessageResponse) GoString() string {
	return s.String()
}

func (s *GetAckInfoOfMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAckInfoOfMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAckInfoOfMessageResponse) GetBody() *GetAckInfoOfMessageResponseBody {
	return s.Body
}

func (s *GetAckInfoOfMessageResponse) SetHeaders(v map[string]*string) *GetAckInfoOfMessageResponse {
	s.Headers = v
	return s
}

func (s *GetAckInfoOfMessageResponse) SetStatusCode(v int32) *GetAckInfoOfMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAckInfoOfMessageResponse) SetBody(v *GetAckInfoOfMessageResponseBody) *GetAckInfoOfMessageResponse {
	s.Body = v
	return s
}

func (s *GetAckInfoOfMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
