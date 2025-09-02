// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensClassificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySensClassificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySensClassificationResponse
	GetStatusCode() *int32
	SetBody(v *QuerySensClassificationResponseBody) *QuerySensClassificationResponse
	GetBody() *QuerySensClassificationResponseBody
}

type QuerySensClassificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySensClassificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySensClassificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySensClassificationResponse) GoString() string {
	return s.String()
}

func (s *QuerySensClassificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySensClassificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySensClassificationResponse) GetBody() *QuerySensClassificationResponseBody {
	return s.Body
}

func (s *QuerySensClassificationResponse) SetHeaders(v map[string]*string) *QuerySensClassificationResponse {
	s.Headers = v
	return s
}

func (s *QuerySensClassificationResponse) SetStatusCode(v int32) *QuerySensClassificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySensClassificationResponse) SetBody(v *QuerySensClassificationResponseBody) *QuerySensClassificationResponse {
	s.Body = v
	return s
}

func (s *QuerySensClassificationResponse) Validate() error {
	return dara.Validate(s)
}
