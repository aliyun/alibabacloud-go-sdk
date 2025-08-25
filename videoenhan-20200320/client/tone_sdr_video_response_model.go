// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToneSdrVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ToneSdrVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ToneSdrVideoResponse
	GetStatusCode() *int32
	SetBody(v *ToneSdrVideoResponseBody) *ToneSdrVideoResponse
	GetBody() *ToneSdrVideoResponseBody
}

type ToneSdrVideoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ToneSdrVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ToneSdrVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s ToneSdrVideoResponse) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ToneSdrVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ToneSdrVideoResponse) GetBody() *ToneSdrVideoResponseBody {
	return s.Body
}

func (s *ToneSdrVideoResponse) SetHeaders(v map[string]*string) *ToneSdrVideoResponse {
	s.Headers = v
	return s
}

func (s *ToneSdrVideoResponse) SetStatusCode(v int32) *ToneSdrVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *ToneSdrVideoResponse) SetBody(v *ToneSdrVideoResponseBody) *ToneSdrVideoResponse {
	s.Body = v
	return s
}

func (s *ToneSdrVideoResponse) Validate() error {
	return dara.Validate(s)
}
