// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyPlayRegionStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDailyPlayRegionStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDailyPlayRegionStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetDailyPlayRegionStatisResponseBody) *GetDailyPlayRegionStatisResponse
	GetBody() *GetDailyPlayRegionStatisResponseBody
}

type GetDailyPlayRegionStatisResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDailyPlayRegionStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDailyPlayRegionStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayRegionStatisResponse) GoString() string {
	return s.String()
}

func (s *GetDailyPlayRegionStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDailyPlayRegionStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDailyPlayRegionStatisResponse) GetBody() *GetDailyPlayRegionStatisResponseBody {
	return s.Body
}

func (s *GetDailyPlayRegionStatisResponse) SetHeaders(v map[string]*string) *GetDailyPlayRegionStatisResponse {
	s.Headers = v
	return s
}

func (s *GetDailyPlayRegionStatisResponse) SetStatusCode(v int32) *GetDailyPlayRegionStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDailyPlayRegionStatisResponse) SetBody(v *GetDailyPlayRegionStatisResponseBody) *GetDailyPlayRegionStatisResponse {
	s.Body = v
	return s
}

func (s *GetDailyPlayRegionStatisResponse) Validate() error {
	return dara.Validate(s)
}
