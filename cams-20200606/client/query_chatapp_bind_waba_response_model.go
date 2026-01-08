// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatappBindWabaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryChatappBindWabaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryChatappBindWabaResponse
	GetStatusCode() *int32
	SetBody(v *QueryChatappBindWabaResponseBody) *QueryChatappBindWabaResponse
	GetBody() *QueryChatappBindWabaResponseBody
}

type QueryChatappBindWabaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryChatappBindWabaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryChatappBindWabaResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappBindWabaResponse) GoString() string {
	return s.String()
}

func (s *QueryChatappBindWabaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryChatappBindWabaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryChatappBindWabaResponse) GetBody() *QueryChatappBindWabaResponseBody {
	return s.Body
}

func (s *QueryChatappBindWabaResponse) SetHeaders(v map[string]*string) *QueryChatappBindWabaResponse {
	s.Headers = v
	return s
}

func (s *QueryChatappBindWabaResponse) SetStatusCode(v int32) *QueryChatappBindWabaResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryChatappBindWabaResponse) SetBody(v *QueryChatappBindWabaResponseBody) *QueryChatappBindWabaResponse {
	s.Body = v
	return s
}

func (s *QueryChatappBindWabaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
