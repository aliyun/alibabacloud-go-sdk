// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuntimeChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuntimeChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuntimeChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetRuntimeChannelResponseBody) *GetRuntimeChannelResponse
	GetBody() *GetRuntimeChannelResponseBody
}

type GetRuntimeChannelResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuntimeChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuntimeChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuntimeChannelResponse) GoString() string {
	return s.String()
}

func (s *GetRuntimeChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuntimeChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuntimeChannelResponse) GetBody() *GetRuntimeChannelResponseBody {
	return s.Body
}

func (s *GetRuntimeChannelResponse) SetHeaders(v map[string]*string) *GetRuntimeChannelResponse {
	s.Headers = v
	return s
}

func (s *GetRuntimeChannelResponse) SetStatusCode(v int32) *GetRuntimeChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuntimeChannelResponse) SetBody(v *GetRuntimeChannelResponseBody) *GetRuntimeChannelResponse {
	s.Body = v
	return s
}

func (s *GetRuntimeChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
