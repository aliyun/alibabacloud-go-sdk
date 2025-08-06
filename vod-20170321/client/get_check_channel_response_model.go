// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckChannelResponseBody) *GetCheckChannelResponse
	GetBody() *GetCheckChannelResponseBody
}

type GetCheckChannelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckChannelResponse) GoString() string {
	return s.String()
}

func (s *GetCheckChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckChannelResponse) GetBody() *GetCheckChannelResponseBody {
	return s.Body
}

func (s *GetCheckChannelResponse) SetHeaders(v map[string]*string) *GetCheckChannelResponse {
	s.Headers = v
	return s
}

func (s *GetCheckChannelResponse) SetStatusCode(v int32) *GetCheckChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckChannelResponse) SetBody(v *GetCheckChannelResponseBody) *GetCheckChannelResponse {
	s.Body = v
	return s
}

func (s *GetCheckChannelResponse) Validate() error {
	return dara.Validate(s)
}
