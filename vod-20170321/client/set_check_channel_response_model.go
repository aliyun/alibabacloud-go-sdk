// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCheckChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCheckChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCheckChannelResponse
	GetStatusCode() *int32
	SetBody(v *SetCheckChannelResponseBody) *SetCheckChannelResponse
	GetBody() *SetCheckChannelResponseBody
}

type SetCheckChannelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCheckChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCheckChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCheckChannelResponse) GoString() string {
	return s.String()
}

func (s *SetCheckChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCheckChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCheckChannelResponse) GetBody() *SetCheckChannelResponseBody {
	return s.Body
}

func (s *SetCheckChannelResponse) SetHeaders(v map[string]*string) *SetCheckChannelResponse {
	s.Headers = v
	return s
}

func (s *SetCheckChannelResponse) SetStatusCode(v int32) *SetCheckChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCheckChannelResponse) SetBody(v *SetCheckChannelResponseBody) *SetCheckChannelResponse {
	s.Body = v
	return s
}

func (s *SetCheckChannelResponse) Validate() error {
	return dara.Validate(s)
}
