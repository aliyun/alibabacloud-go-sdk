// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeSQLLineageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeSQLLineageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeSQLLineageResponse
	GetStatusCode() *int32
	SetBody(v *AnalyzeSQLLineageResponseBody) *AnalyzeSQLLineageResponse
	GetBody() *AnalyzeSQLLineageResponseBody
}

type AnalyzeSQLLineageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeSQLLineageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeSQLLineageResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeSQLLineageResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeSQLLineageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeSQLLineageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeSQLLineageResponse) GetBody() *AnalyzeSQLLineageResponseBody {
	return s.Body
}

func (s *AnalyzeSQLLineageResponse) SetHeaders(v map[string]*string) *AnalyzeSQLLineageResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeSQLLineageResponse) SetStatusCode(v int32) *AnalyzeSQLLineageResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeSQLLineageResponse) SetBody(v *AnalyzeSQLLineageResponseBody) *AnalyzeSQLLineageResponse {
	s.Body = v
	return s
}

func (s *AnalyzeSQLLineageResponse) Validate() error {
	return dara.Validate(s)
}
