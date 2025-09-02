// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySensNodeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySensNodeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySensNodeInfoResponse
	GetStatusCode() *int32
	SetBody(v *QuerySensNodeInfoResponseBody) *QuerySensNodeInfoResponse
	GetBody() *QuerySensNodeInfoResponseBody
}

type QuerySensNodeInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySensNodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySensNodeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySensNodeInfoResponse) GoString() string {
	return s.String()
}

func (s *QuerySensNodeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySensNodeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySensNodeInfoResponse) GetBody() *QuerySensNodeInfoResponseBody {
	return s.Body
}

func (s *QuerySensNodeInfoResponse) SetHeaders(v map[string]*string) *QuerySensNodeInfoResponse {
	s.Headers = v
	return s
}

func (s *QuerySensNodeInfoResponse) SetStatusCode(v int32) *QuerySensNodeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySensNodeInfoResponse) SetBody(v *QuerySensNodeInfoResponseBody) *QuerySensNodeInfoResponse {
	s.Body = v
	return s
}

func (s *QuerySensNodeInfoResponse) Validate() error {
	return dara.Validate(s)
}
