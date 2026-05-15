// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAiOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAiOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAiOutboundTaskResponseBody) *CreateAiOutboundTaskResponse
	GetBody() *CreateAiOutboundTaskResponseBody
}

type CreateAiOutboundTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAiOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAiOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAiOutboundTaskResponse) GetBody() *CreateAiOutboundTaskResponseBody {
	return s.Body
}

func (s *CreateAiOutboundTaskResponse) SetHeaders(v map[string]*string) *CreateAiOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAiOutboundTaskResponse) SetStatusCode(v int32) *CreateAiOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiOutboundTaskResponse) SetBody(v *CreateAiOutboundTaskResponseBody) *CreateAiOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *CreateAiOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
