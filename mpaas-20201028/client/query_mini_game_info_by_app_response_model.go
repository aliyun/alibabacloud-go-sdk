// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMiniGameInfoByAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMiniGameInfoByAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMiniGameInfoByAppResponse
	GetStatusCode() *int32
	SetBody(v *QueryMiniGameInfoByAppResponseBody) *QueryMiniGameInfoByAppResponse
	GetBody() *QueryMiniGameInfoByAppResponseBody
}

type QueryMiniGameInfoByAppResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMiniGameInfoByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMiniGameInfoByAppResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMiniGameInfoByAppResponse) GoString() string {
	return s.String()
}

func (s *QueryMiniGameInfoByAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMiniGameInfoByAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMiniGameInfoByAppResponse) GetBody() *QueryMiniGameInfoByAppResponseBody {
	return s.Body
}

func (s *QueryMiniGameInfoByAppResponse) SetHeaders(v map[string]*string) *QueryMiniGameInfoByAppResponse {
	s.Headers = v
	return s
}

func (s *QueryMiniGameInfoByAppResponse) SetStatusCode(v int32) *QueryMiniGameInfoByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMiniGameInfoByAppResponse) SetBody(v *QueryMiniGameInfoByAppResponseBody) *QueryMiniGameInfoByAppResponse {
	s.Body = v
	return s
}

func (s *QueryMiniGameInfoByAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
