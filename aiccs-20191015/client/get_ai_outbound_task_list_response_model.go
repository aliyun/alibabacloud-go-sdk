// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiOutboundTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiOutboundTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiOutboundTaskListResponse
	GetStatusCode() *int32
	SetBody(v *GetAiOutboundTaskListResponseBody) *GetAiOutboundTaskListResponse
	GetBody() *GetAiOutboundTaskListResponseBody
}

type GetAiOutboundTaskListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiOutboundTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiOutboundTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiOutboundTaskListResponse) GoString() string {
	return s.String()
}

func (s *GetAiOutboundTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiOutboundTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiOutboundTaskListResponse) GetBody() *GetAiOutboundTaskListResponseBody {
	return s.Body
}

func (s *GetAiOutboundTaskListResponse) SetHeaders(v map[string]*string) *GetAiOutboundTaskListResponse {
	s.Headers = v
	return s
}

func (s *GetAiOutboundTaskListResponse) SetStatusCode(v int32) *GetAiOutboundTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiOutboundTaskListResponse) SetBody(v *GetAiOutboundTaskListResponseBody) *GetAiOutboundTaskListResponse {
	s.Body = v
	return s
}

func (s *GetAiOutboundTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
