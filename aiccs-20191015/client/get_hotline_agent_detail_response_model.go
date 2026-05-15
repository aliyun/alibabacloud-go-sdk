// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineAgentDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineAgentDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineAgentDetailResponseBody) *GetHotlineAgentDetailResponse
	GetBody() *GetHotlineAgentDetailResponseBody
}

type GetHotlineAgentDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineAgentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineAgentDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineAgentDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineAgentDetailResponse) GetBody() *GetHotlineAgentDetailResponseBody {
	return s.Body
}

func (s *GetHotlineAgentDetailResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailResponse) SetStatusCode(v int32) *GetHotlineAgentDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponse) SetBody(v *GetHotlineAgentDetailResponseBody) *GetHotlineAgentDetailResponse {
	s.Body = v
	return s
}

func (s *GetHotlineAgentDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
