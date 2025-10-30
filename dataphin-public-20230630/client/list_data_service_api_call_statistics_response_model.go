// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiCallStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataServiceApiCallStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataServiceApiCallStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataServiceApiCallStatisticsResponseBody) *ListDataServiceApiCallStatisticsResponse
	GetBody() *ListDataServiceApiCallStatisticsResponseBody
}

type ListDataServiceApiCallStatisticsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataServiceApiCallStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataServiceApiCallStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiCallStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiCallStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataServiceApiCallStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataServiceApiCallStatisticsResponse) GetBody() *ListDataServiceApiCallStatisticsResponseBody {
	return s.Body
}

func (s *ListDataServiceApiCallStatisticsResponse) SetHeaders(v map[string]*string) *ListDataServiceApiCallStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponse) SetStatusCode(v int32) *ListDataServiceApiCallStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponse) SetBody(v *ListDataServiceApiCallStatisticsResponseBody) *ListDataServiceApiCallStatisticsResponse {
	s.Body = v
	return s
}

func (s *ListDataServiceApiCallStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
