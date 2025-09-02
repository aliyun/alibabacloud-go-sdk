// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQuerySensResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgQuerySensResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgQuerySensResultResponse
	GetStatusCode() *int32
	SetBody(v *DsgQuerySensResultResponseBody) *DsgQuerySensResultResponse
	GetBody() *DsgQuerySensResultResponseBody
}

type DsgQuerySensResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgQuerySensResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgQuerySensResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgQuerySensResultResponse) GoString() string {
	return s.String()
}

func (s *DsgQuerySensResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgQuerySensResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgQuerySensResultResponse) GetBody() *DsgQuerySensResultResponseBody {
	return s.Body
}

func (s *DsgQuerySensResultResponse) SetHeaders(v map[string]*string) *DsgQuerySensResultResponse {
	s.Headers = v
	return s
}

func (s *DsgQuerySensResultResponse) SetStatusCode(v int32) *DsgQuerySensResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgQuerySensResultResponse) SetBody(v *DsgQuerySensResultResponseBody) *DsgQuerySensResultResponse {
	s.Body = v
	return s
}

func (s *DsgQuerySensResultResponse) Validate() error {
	return dara.Validate(s)
}
