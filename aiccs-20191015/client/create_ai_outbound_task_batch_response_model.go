// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAiOutboundTaskBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAiOutboundTaskBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAiOutboundTaskBatchResponse
	GetStatusCode() *int32
	SetBody(v *CreateAiOutboundTaskBatchResponseBody) *CreateAiOutboundTaskBatchResponse
	GetBody() *CreateAiOutboundTaskBatchResponseBody
}

type CreateAiOutboundTaskBatchResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAiOutboundTaskBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAiOutboundTaskBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAiOutboundTaskBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateAiOutboundTaskBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAiOutboundTaskBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAiOutboundTaskBatchResponse) GetBody() *CreateAiOutboundTaskBatchResponseBody {
	return s.Body
}

func (s *CreateAiOutboundTaskBatchResponse) SetHeaders(v map[string]*string) *CreateAiOutboundTaskBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateAiOutboundTaskBatchResponse) SetStatusCode(v int32) *CreateAiOutboundTaskBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAiOutboundTaskBatchResponse) SetBody(v *CreateAiOutboundTaskBatchResponseBody) *CreateAiOutboundTaskBatchResponse {
	s.Body = v
	return s
}

func (s *CreateAiOutboundTaskBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
