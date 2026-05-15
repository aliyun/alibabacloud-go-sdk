// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineRuntimeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineRuntimeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineRuntimeInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineRuntimeInfoResponseBody) *GetHotlineRuntimeInfoResponse
	GetBody() *GetHotlineRuntimeInfoResponseBody
}

type GetHotlineRuntimeInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineRuntimeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineRuntimeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineRuntimeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineRuntimeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineRuntimeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineRuntimeInfoResponse) GetBody() *GetHotlineRuntimeInfoResponseBody {
	return s.Body
}

func (s *GetHotlineRuntimeInfoResponse) SetHeaders(v map[string]*string) *GetHotlineRuntimeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineRuntimeInfoResponse) SetStatusCode(v int32) *GetHotlineRuntimeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponse) SetBody(v *GetHotlineRuntimeInfoResponseBody) *GetHotlineRuntimeInfoResponse {
	s.Body = v
	return s
}

func (s *GetHotlineRuntimeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
