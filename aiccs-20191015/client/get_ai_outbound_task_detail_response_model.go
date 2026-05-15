// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiOutboundTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiOutboundTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetAiOutboundTaskDetailResponseBody) *GetAiOutboundTaskDetailResponse
	GetBody() *GetAiOutboundTaskDetailResponseBody
}

type GetAiOutboundTaskDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiOutboundTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiOutboundTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiOutboundTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiOutboundTaskDetailResponse) GetBody() *GetAiOutboundTaskDetailResponseBody {
	return s.Body
}

func (s *GetAiOutboundTaskDetailResponse) SetHeaders(v map[string]*string) *GetAiOutboundTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetAiOutboundTaskDetailResponse) SetStatusCode(v int32) *GetAiOutboundTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiOutboundTaskDetailResponse) SetBody(v *GetAiOutboundTaskDetailResponseBody) *GetAiOutboundTaskDetailResponse {
	s.Body = v
	return s
}

func (s *GetAiOutboundTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
