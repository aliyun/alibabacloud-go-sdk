// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommercialUsageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCommercialUsageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCommercialUsageResponse
	GetStatusCode() *int32
	SetBody(v *QueryCommercialUsageResponseBody) *QueryCommercialUsageResponse
	GetBody() *QueryCommercialUsageResponseBody
}

type QueryCommercialUsageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCommercialUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCommercialUsageResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCommercialUsageResponse) GoString() string {
	return s.String()
}

func (s *QueryCommercialUsageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCommercialUsageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCommercialUsageResponse) GetBody() *QueryCommercialUsageResponseBody {
	return s.Body
}

func (s *QueryCommercialUsageResponse) SetHeaders(v map[string]*string) *QueryCommercialUsageResponse {
	s.Headers = v
	return s
}

func (s *QueryCommercialUsageResponse) SetStatusCode(v int32) *QueryCommercialUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCommercialUsageResponse) SetBody(v *QueryCommercialUsageResponseBody) *QueryCommercialUsageResponse {
	s.Body = v
	return s
}

func (s *QueryCommercialUsageResponse) Validate() error {
	return dara.Validate(s)
}
