// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDynamicSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDynamicSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDynamicSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDynamicSettingsResponseBody) *DescribeDynamicSettingsResponse
	GetBody() *DescribeDynamicSettingsResponseBody
}

type DescribeDynamicSettingsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDynamicSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDynamicSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDynamicSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDynamicSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDynamicSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDynamicSettingsResponse) GetBody() *DescribeDynamicSettingsResponseBody {
	return s.Body
}

func (s *DescribeDynamicSettingsResponse) SetHeaders(v map[string]*string) *DescribeDynamicSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDynamicSettingsResponse) SetStatusCode(v int32) *DescribeDynamicSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDynamicSettingsResponse) SetBody(v *DescribeDynamicSettingsResponseBody) *DescribeDynamicSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeDynamicSettingsResponse) Validate() error {
	return dara.Validate(s)
}
