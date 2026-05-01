// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRuntimeChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveRuntimeChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveRuntimeChannelResponse
	GetStatusCode() *int32
	SetBody(v *RemoveRuntimeChannelResponseBody) *RemoveRuntimeChannelResponse
	GetBody() *RemoveRuntimeChannelResponseBody
}

type RemoveRuntimeChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRuntimeChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRuntimeChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveRuntimeChannelResponse) GoString() string {
	return s.String()
}

func (s *RemoveRuntimeChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveRuntimeChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveRuntimeChannelResponse) GetBody() *RemoveRuntimeChannelResponseBody {
	return s.Body
}

func (s *RemoveRuntimeChannelResponse) SetHeaders(v map[string]*string) *RemoveRuntimeChannelResponse {
	s.Headers = v
	return s
}

func (s *RemoveRuntimeChannelResponse) SetStatusCode(v int32) *RemoveRuntimeChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRuntimeChannelResponse) SetBody(v *RemoveRuntimeChannelResponseBody) *RemoveRuntimeChannelResponse {
	s.Body = v
	return s
}

func (s *RemoveRuntimeChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
