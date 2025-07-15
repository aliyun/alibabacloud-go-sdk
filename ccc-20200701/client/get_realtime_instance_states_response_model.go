// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeInstanceStatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealtimeInstanceStatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealtimeInstanceStatesResponse
	GetStatusCode() *int32
	SetBody(v *GetRealtimeInstanceStatesResponseBody) *GetRealtimeInstanceStatesResponse
	GetBody() *GetRealtimeInstanceStatesResponseBody
}

type GetRealtimeInstanceStatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeInstanceStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeInstanceStatesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeInstanceStatesResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeInstanceStatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealtimeInstanceStatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealtimeInstanceStatesResponse) GetBody() *GetRealtimeInstanceStatesResponseBody {
	return s.Body
}

func (s *GetRealtimeInstanceStatesResponse) SetHeaders(v map[string]*string) *GetRealtimeInstanceStatesResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeInstanceStatesResponse) SetStatusCode(v int32) *GetRealtimeInstanceStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeInstanceStatesResponse) SetBody(v *GetRealtimeInstanceStatesResponseBody) *GetRealtimeInstanceStatesResponse {
	s.Body = v
	return s
}

func (s *GetRealtimeInstanceStatesResponse) Validate() error {
	return dara.Validate(s)
}
