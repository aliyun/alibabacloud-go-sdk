// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappBindWabaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappBindWabaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappBindWabaResponse
	GetStatusCode() *int32
	SetBody(v *ChatappBindWabaResponseBody) *ChatappBindWabaResponse
	GetBody() *ChatappBindWabaResponseBody
}

type ChatappBindWabaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappBindWabaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappBindWabaResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappBindWabaResponse) GoString() string {
	return s.String()
}

func (s *ChatappBindWabaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappBindWabaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappBindWabaResponse) GetBody() *ChatappBindWabaResponseBody {
	return s.Body
}

func (s *ChatappBindWabaResponse) SetHeaders(v map[string]*string) *ChatappBindWabaResponse {
	s.Headers = v
	return s
}

func (s *ChatappBindWabaResponse) SetStatusCode(v int32) *ChatappBindWabaResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappBindWabaResponse) SetBody(v *ChatappBindWabaResponseBody) *ChatappBindWabaResponse {
	s.Body = v
	return s
}

func (s *ChatappBindWabaResponse) Validate() error {
	return dara.Validate(s)
}
