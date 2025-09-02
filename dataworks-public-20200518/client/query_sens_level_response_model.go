// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensLevelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySensLevelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySensLevelResponse
	GetStatusCode() *int32
	SetBody(v *QuerySensLevelResponseBody) *QuerySensLevelResponse
	GetBody() *QuerySensLevelResponseBody
}

type QuerySensLevelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySensLevelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySensLevelResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySensLevelResponse) GoString() string {
	return s.String()
}

func (s *QuerySensLevelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySensLevelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySensLevelResponse) GetBody() *QuerySensLevelResponseBody {
	return s.Body
}

func (s *QuerySensLevelResponse) SetHeaders(v map[string]*string) *QuerySensLevelResponse {
	s.Headers = v
	return s
}

func (s *QuerySensLevelResponse) SetStatusCode(v int32) *QuerySensLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySensLevelResponse) SetBody(v *QuerySensLevelResponseBody) *QuerySensLevelResponse {
	s.Body = v
	return s
}

func (s *QuerySensLevelResponse) Validate() error {
	return dara.Validate(s)
}
