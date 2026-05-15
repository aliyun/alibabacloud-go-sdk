// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskExecDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiOutboundTaskExecDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiOutboundTaskExecDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAiOutboundTaskExecDetailResponseBody) *GetAiOutboundTaskExecDetailResponse
	GetBody() *GetAiOutboundTaskExecDetailResponseBody
}

type GetAiOutboundTaskExecDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiOutboundTaskExecDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiOutboundTaskExecDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskExecDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskExecDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiOutboundTaskExecDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiOutboundTaskExecDetailResponse) GetBody() *GetAiOutboundTaskExecDetailResponseBody {
	return s.Body
}

func (s *GetAiOutboundTaskExecDetailResponse) SetHeaders(v map[string]*string) *GetAiOutboundTaskExecDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponse) SetStatusCode(v int32) *GetAiOutboundTaskExecDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponse) SetBody(v *GetAiOutboundTaskExecDetailResponseBody) *GetAiOutboundTaskExecDetailResponse {
	s.Body = v
	return s
}

func (s *GetAiOutboundTaskExecDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
