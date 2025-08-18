// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafBotAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWafBotAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWafBotAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetWafBotAppKeyResponseBody) *GetWafBotAppKeyResponse
	GetBody() *GetWafBotAppKeyResponseBody
}

type GetWafBotAppKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafBotAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafBotAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWafBotAppKeyResponse) GoString() string {
	return s.String()
}

func (s *GetWafBotAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWafBotAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWafBotAppKeyResponse) GetBody() *GetWafBotAppKeyResponseBody {
	return s.Body
}

func (s *GetWafBotAppKeyResponse) SetHeaders(v map[string]*string) *GetWafBotAppKeyResponse {
	s.Headers = v
	return s
}

func (s *GetWafBotAppKeyResponse) SetStatusCode(v int32) *GetWafBotAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafBotAppKeyResponse) SetBody(v *GetWafBotAppKeyResponseBody) *GetWafBotAppKeyResponse {
	s.Body = v
	return s
}

func (s *GetWafBotAppKeyResponse) Validate() error {
	return dara.Validate(s)
}
