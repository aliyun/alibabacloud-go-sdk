// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTrademarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmsTrademarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmsTrademarkResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmsTrademarkResponseBody) *QuerySmsTrademarkResponse
	GetBody() *QuerySmsTrademarkResponseBody
}

type QuerySmsTrademarkResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmsTrademarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmsTrademarkResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTrademarkResponse) GoString() string {
	return s.String()
}

func (s *QuerySmsTrademarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmsTrademarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmsTrademarkResponse) GetBody() *QuerySmsTrademarkResponseBody {
	return s.Body
}

func (s *QuerySmsTrademarkResponse) SetHeaders(v map[string]*string) *QuerySmsTrademarkResponse {
	s.Headers = v
	return s
}

func (s *QuerySmsTrademarkResponse) SetStatusCode(v int32) *QuerySmsTrademarkResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmsTrademarkResponse) SetBody(v *QuerySmsTrademarkResponseBody) *QuerySmsTrademarkResponse {
	s.Body = v
	return s
}

func (s *QuerySmsTrademarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
