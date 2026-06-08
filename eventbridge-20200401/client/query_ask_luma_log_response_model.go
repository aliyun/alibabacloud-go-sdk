// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAskLumaLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAskLumaLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAskLumaLogResponse
	GetStatusCode() *int32
	SetBody(v *QueryAskLumaLogResponseBody) *QueryAskLumaLogResponse
	GetBody() *QueryAskLumaLogResponseBody
}

type QueryAskLumaLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAskLumaLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAskLumaLogResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAskLumaLogResponse) GoString() string {
	return s.String()
}

func (s *QueryAskLumaLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAskLumaLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAskLumaLogResponse) GetBody() *QueryAskLumaLogResponseBody {
	return s.Body
}

func (s *QueryAskLumaLogResponse) SetHeaders(v map[string]*string) *QueryAskLumaLogResponse {
	s.Headers = v
	return s
}

func (s *QueryAskLumaLogResponse) SetStatusCode(v int32) *QueryAskLumaLogResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAskLumaLogResponse) SetBody(v *QueryAskLumaLogResponseBody) *QueryAskLumaLogResponse {
	s.Body = v
	return s
}

func (s *QueryAskLumaLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
