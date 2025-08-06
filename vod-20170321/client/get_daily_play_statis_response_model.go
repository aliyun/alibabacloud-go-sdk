// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyPlayStatisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDailyPlayStatisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDailyPlayStatisResponse
	GetStatusCode() *int32
	SetBody(v *GetDailyPlayStatisResponseBody) *GetDailyPlayStatisResponse
	GetBody() *GetDailyPlayStatisResponseBody
}

type GetDailyPlayStatisResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDailyPlayStatisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDailyPlayStatisResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDailyPlayStatisResponse) GoString() string {
	return s.String()
}

func (s *GetDailyPlayStatisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDailyPlayStatisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDailyPlayStatisResponse) GetBody() *GetDailyPlayStatisResponseBody {
	return s.Body
}

func (s *GetDailyPlayStatisResponse) SetHeaders(v map[string]*string) *GetDailyPlayStatisResponse {
	s.Headers = v
	return s
}

func (s *GetDailyPlayStatisResponse) SetStatusCode(v int32) *GetDailyPlayStatisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDailyPlayStatisResponse) SetBody(v *GetDailyPlayStatisResponseBody) *GetDailyPlayStatisResponse {
	s.Body = v
	return s
}

func (s *GetDailyPlayStatisResponse) Validate() error {
	return dara.Validate(s)
}
