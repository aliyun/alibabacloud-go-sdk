// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlashbackRecycleBinTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlashbackRecycleBinTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlashbackRecycleBinTableResponse
	GetStatusCode() *int32
	SetBody(v *FlashbackRecycleBinTableResponseBody) *FlashbackRecycleBinTableResponse
	GetBody() *FlashbackRecycleBinTableResponseBody
}

type FlashbackRecycleBinTableResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlashbackRecycleBinTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlashbackRecycleBinTableResponse) String() string {
	return dara.Prettify(s)
}

func (s FlashbackRecycleBinTableResponse) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlashbackRecycleBinTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlashbackRecycleBinTableResponse) GetBody() *FlashbackRecycleBinTableResponseBody {
	return s.Body
}

func (s *FlashbackRecycleBinTableResponse) SetHeaders(v map[string]*string) *FlashbackRecycleBinTableResponse {
	s.Headers = v
	return s
}

func (s *FlashbackRecycleBinTableResponse) SetStatusCode(v int32) *FlashbackRecycleBinTableResponse {
	s.StatusCode = &v
	return s
}

func (s *FlashbackRecycleBinTableResponse) SetBody(v *FlashbackRecycleBinTableResponseBody) *FlashbackRecycleBinTableResponse {
	s.Body = v
	return s
}

func (s *FlashbackRecycleBinTableResponse) Validate() error {
	return dara.Validate(s)
}
