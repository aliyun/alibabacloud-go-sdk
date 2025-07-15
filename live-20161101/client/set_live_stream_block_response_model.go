// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamBlockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveStreamBlockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveStreamBlockResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveStreamBlockResponseBody) *SetLiveStreamBlockResponse
	GetBody() *SetLiveStreamBlockResponseBody
}

type SetLiveStreamBlockResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveStreamBlockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveStreamBlockResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamBlockResponse) GoString() string {
	return s.String()
}

func (s *SetLiveStreamBlockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveStreamBlockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveStreamBlockResponse) GetBody() *SetLiveStreamBlockResponseBody {
	return s.Body
}

func (s *SetLiveStreamBlockResponse) SetHeaders(v map[string]*string) *SetLiveStreamBlockResponse {
	s.Headers = v
	return s
}

func (s *SetLiveStreamBlockResponse) SetStatusCode(v int32) *SetLiveStreamBlockResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveStreamBlockResponse) SetBody(v *SetLiveStreamBlockResponseBody) *SetLiveStreamBlockResponse {
	s.Body = v
	return s
}

func (s *SetLiveStreamBlockResponse) Validate() error {
	return dara.Validate(s)
}
