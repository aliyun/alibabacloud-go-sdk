// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityScanRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataQualityScanRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataQualityScanRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataQualityScanRunsResponseBody) *ListDataQualityScanRunsResponse
	GetBody() *ListDataQualityScanRunsResponseBody
}

type ListDataQualityScanRunsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityScanRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityScanRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScanRunsResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityScanRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataQualityScanRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataQualityScanRunsResponse) GetBody() *ListDataQualityScanRunsResponseBody {
	return s.Body
}

func (s *ListDataQualityScanRunsResponse) SetHeaders(v map[string]*string) *ListDataQualityScanRunsResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityScanRunsResponse) SetStatusCode(v int32) *ListDataQualityScanRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityScanRunsResponse) SetBody(v *ListDataQualityScanRunsResponseBody) *ListDataQualityScanRunsResponse {
	s.Body = v
	return s
}

func (s *ListDataQualityScanRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
