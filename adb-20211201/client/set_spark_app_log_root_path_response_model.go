// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSparkAppLogRootPathResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSparkAppLogRootPathResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSparkAppLogRootPathResponse
	GetStatusCode() *int32
	SetBody(v *SetSparkAppLogRootPathResponseBody) *SetSparkAppLogRootPathResponse
	GetBody() *SetSparkAppLogRootPathResponseBody
}

type SetSparkAppLogRootPathResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSparkAppLogRootPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSparkAppLogRootPathResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSparkAppLogRootPathResponse) GoString() string {
	return s.String()
}

func (s *SetSparkAppLogRootPathResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSparkAppLogRootPathResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSparkAppLogRootPathResponse) GetBody() *SetSparkAppLogRootPathResponseBody {
	return s.Body
}

func (s *SetSparkAppLogRootPathResponse) SetHeaders(v map[string]*string) *SetSparkAppLogRootPathResponse {
	s.Headers = v
	return s
}

func (s *SetSparkAppLogRootPathResponse) SetStatusCode(v int32) *SetSparkAppLogRootPathResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSparkAppLogRootPathResponse) SetBody(v *SetSparkAppLogRootPathResponseBody) *SetSparkAppLogRootPathResponse {
	s.Body = v
	return s
}

func (s *SetSparkAppLogRootPathResponse) Validate() error {
	return dara.Validate(s)
}
