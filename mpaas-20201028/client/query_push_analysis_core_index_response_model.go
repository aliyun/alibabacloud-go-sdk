// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisCoreIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPushAnalysisCoreIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPushAnalysisCoreIndexResponse
	GetStatusCode() *int32
	SetBody(v *QueryPushAnalysisCoreIndexResponseBody) *QueryPushAnalysisCoreIndexResponse
	GetBody() *QueryPushAnalysisCoreIndexResponseBody
}

type QueryPushAnalysisCoreIndexResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPushAnalysisCoreIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPushAnalysisCoreIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisCoreIndexResponse) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisCoreIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPushAnalysisCoreIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPushAnalysisCoreIndexResponse) GetBody() *QueryPushAnalysisCoreIndexResponseBody {
	return s.Body
}

func (s *QueryPushAnalysisCoreIndexResponse) SetHeaders(v map[string]*string) *QueryPushAnalysisCoreIndexResponse {
	s.Headers = v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponse) SetStatusCode(v int32) *QueryPushAnalysisCoreIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponse) SetBody(v *QueryPushAnalysisCoreIndexResponseBody) *QueryPushAnalysisCoreIndexResponse {
	s.Body = v
	return s
}

func (s *QueryPushAnalysisCoreIndexResponse) Validate() error {
	return dara.Validate(s)
}
