// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeMiniPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMcubeMiniPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMcubeMiniPackageResponse
	GetStatusCode() *int32
	SetBody(v *QueryMcubeMiniPackageResponseBody) *QueryMcubeMiniPackageResponse
	GetBody() *QueryMcubeMiniPackageResponseBody
}

type QueryMcubeMiniPackageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMcubeMiniPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMcubeMiniPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniPackageResponse) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMcubeMiniPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMcubeMiniPackageResponse) GetBody() *QueryMcubeMiniPackageResponseBody {
	return s.Body
}

func (s *QueryMcubeMiniPackageResponse) SetHeaders(v map[string]*string) *QueryMcubeMiniPackageResponse {
	s.Headers = v
	return s
}

func (s *QueryMcubeMiniPackageResponse) SetStatusCode(v int32) *QueryMcubeMiniPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMcubeMiniPackageResponse) SetBody(v *QueryMcubeMiniPackageResponseBody) *QueryMcubeMiniPackageResponse {
	s.Body = v
	return s
}

func (s *QueryMcubeMiniPackageResponse) Validate() error {
	return dara.Validate(s)
}
