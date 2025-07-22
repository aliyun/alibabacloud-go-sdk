// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceGroupStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceGroupStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceGroupStatisticsResponseBody) *GetResourceGroupStatisticsResponse
	GetBody() *GetResourceGroupStatisticsResponseBody
}

type GetResourceGroupStatisticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceGroupStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceGroupStatisticsResponse) GetBody() *GetResourceGroupStatisticsResponseBody {
	return s.Body
}

func (s *GetResourceGroupStatisticsResponse) SetHeaders(v map[string]*string) *GetResourceGroupStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupStatisticsResponse) SetStatusCode(v int32) *GetResourceGroupStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupStatisticsResponse) SetBody(v *GetResourceGroupStatisticsResponseBody) *GetResourceGroupStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetResourceGroupStatisticsResponse) Validate() error {
	return dara.Validate(s)
}
