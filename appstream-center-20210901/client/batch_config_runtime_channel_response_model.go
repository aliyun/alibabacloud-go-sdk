// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchConfigRuntimeChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchConfigRuntimeChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchConfigRuntimeChannelResponse
	GetStatusCode() *int32
	SetBody(v *BatchConfigRuntimeChannelResponseBody) *BatchConfigRuntimeChannelResponse
	GetBody() *BatchConfigRuntimeChannelResponseBody
}

type BatchConfigRuntimeChannelResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchConfigRuntimeChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchConfigRuntimeChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchConfigRuntimeChannelResponse) GoString() string {
	return s.String()
}

func (s *BatchConfigRuntimeChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchConfigRuntimeChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchConfigRuntimeChannelResponse) GetBody() *BatchConfigRuntimeChannelResponseBody {
	return s.Body
}

func (s *BatchConfigRuntimeChannelResponse) SetHeaders(v map[string]*string) *BatchConfigRuntimeChannelResponse {
	s.Headers = v
	return s
}

func (s *BatchConfigRuntimeChannelResponse) SetStatusCode(v int32) *BatchConfigRuntimeChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchConfigRuntimeChannelResponse) SetBody(v *BatchConfigRuntimeChannelResponseBody) *BatchConfigRuntimeChannelResponse {
	s.Body = v
	return s
}

func (s *BatchConfigRuntimeChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
