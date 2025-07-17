// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnCallSchedulesDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOnCallSchedulesDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOnCallSchedulesDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetOnCallSchedulesDetailResponseBody) *GetOnCallSchedulesDetailResponse
	GetBody() *GetOnCallSchedulesDetailResponseBody
}

type GetOnCallSchedulesDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOnCallSchedulesDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOnCallSchedulesDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOnCallSchedulesDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOnCallSchedulesDetailResponse) GetBody() *GetOnCallSchedulesDetailResponseBody {
	return s.Body
}

func (s *GetOnCallSchedulesDetailResponse) SetHeaders(v map[string]*string) *GetOnCallSchedulesDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOnCallSchedulesDetailResponse) SetStatusCode(v int32) *GetOnCallSchedulesDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponse) SetBody(v *GetOnCallSchedulesDetailResponseBody) *GetOnCallSchedulesDetailResponse {
	s.Body = v
	return s
}

func (s *GetOnCallSchedulesDetailResponse) Validate() error {
	return dara.Validate(s)
}
