// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterGroupChangeLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterGroupChangeLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLakebaseMountParameterGroupChangeLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLakebaseMountParameterGroupChangeLogsResponseBody) *DescribeLakebaseMountParameterGroupChangeLogsResponse
	GetBody() *DescribeLakebaseMountParameterGroupChangeLogsResponseBody
}

type DescribeLakebaseMountParameterGroupChangeLogsResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLakebaseMountParameterGroupChangeLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLakebaseMountParameterGroupChangeLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterGroupChangeLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponse) GetBody() *DescribeLakebaseMountParameterGroupChangeLogsResponseBody {
	return s.Body
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponse) SetHeaders(v map[string]*string) *DescribeLakebaseMountParameterGroupChangeLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponse) SetStatusCode(v int32) *DescribeLakebaseMountParameterGroupChangeLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponse) SetBody(v *DescribeLakebaseMountParameterGroupChangeLogsResponseBody) *DescribeLakebaseMountParameterGroupChangeLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeLakebaseMountParameterGroupChangeLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
