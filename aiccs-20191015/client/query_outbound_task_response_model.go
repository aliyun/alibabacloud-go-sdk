// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *QueryOutboundTaskResponseBody) *QueryOutboundTaskResponse
	GetBody() *QueryOutboundTaskResponseBody
}

type QueryOutboundTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOutboundTaskResponse) GetBody() *QueryOutboundTaskResponseBody {
	return s.Body
}

func (s *QueryOutboundTaskResponse) SetHeaders(v map[string]*string) *QueryOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryOutboundTaskResponse) SetStatusCode(v int32) *QueryOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOutboundTaskResponse) SetBody(v *QueryOutboundTaskResponseBody) *QueryOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *QueryOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
