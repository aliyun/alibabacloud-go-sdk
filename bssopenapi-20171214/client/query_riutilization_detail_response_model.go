// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRIUtilizationDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRIUtilizationDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRIUtilizationDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryRIUtilizationDetailResponseBody) *QueryRIUtilizationDetailResponse
	GetBody() *QueryRIUtilizationDetailResponseBody
}

type QueryRIUtilizationDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRIUtilizationDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRIUtilizationDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRIUtilizationDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRIUtilizationDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRIUtilizationDetailResponse) GetBody() *QueryRIUtilizationDetailResponseBody {
	return s.Body
}

func (s *QueryRIUtilizationDetailResponse) SetHeaders(v map[string]*string) *QueryRIUtilizationDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRIUtilizationDetailResponse) SetStatusCode(v int32) *QueryRIUtilizationDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRIUtilizationDetailResponse) SetBody(v *QueryRIUtilizationDetailResponseBody) *QueryRIUtilizationDetailResponse {
	s.Body = v
	return s
}

func (s *QueryRIUtilizationDetailResponse) Validate() error {
	return dara.Validate(s)
}
