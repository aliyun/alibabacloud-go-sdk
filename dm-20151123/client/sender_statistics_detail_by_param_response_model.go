// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSenderStatisticsDetailByParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SenderStatisticsDetailByParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SenderStatisticsDetailByParamResponse
	GetStatusCode() *int32
	SetBody(v *SenderStatisticsDetailByParamResponseBody) *SenderStatisticsDetailByParamResponse
	GetBody() *SenderStatisticsDetailByParamResponseBody
}

type SenderStatisticsDetailByParamResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SenderStatisticsDetailByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SenderStatisticsDetailByParamResponse) String() string {
	return dara.Prettify(s)
}

func (s SenderStatisticsDetailByParamResponse) GoString() string {
	return s.String()
}

func (s *SenderStatisticsDetailByParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SenderStatisticsDetailByParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SenderStatisticsDetailByParamResponse) GetBody() *SenderStatisticsDetailByParamResponseBody {
	return s.Body
}

func (s *SenderStatisticsDetailByParamResponse) SetHeaders(v map[string]*string) *SenderStatisticsDetailByParamResponse {
	s.Headers = v
	return s
}

func (s *SenderStatisticsDetailByParamResponse) SetStatusCode(v int32) *SenderStatisticsDetailByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *SenderStatisticsDetailByParamResponse) SetBody(v *SenderStatisticsDetailByParamResponseBody) *SenderStatisticsDetailByParamResponse {
	s.Body = v
	return s
}

func (s *SenderStatisticsDetailByParamResponse) Validate() error {
	return dara.Validate(s)
}
