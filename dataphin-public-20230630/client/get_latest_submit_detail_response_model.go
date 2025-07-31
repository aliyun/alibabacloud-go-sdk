// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestSubmitDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLatestSubmitDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLatestSubmitDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetLatestSubmitDetailResponseBody) *GetLatestSubmitDetailResponse
	GetBody() *GetLatestSubmitDetailResponseBody
}

type GetLatestSubmitDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLatestSubmitDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLatestSubmitDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailResponse) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLatestSubmitDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLatestSubmitDetailResponse) GetBody() *GetLatestSubmitDetailResponseBody {
	return s.Body
}

func (s *GetLatestSubmitDetailResponse) SetHeaders(v map[string]*string) *GetLatestSubmitDetailResponse {
	s.Headers = v
	return s
}

func (s *GetLatestSubmitDetailResponse) SetStatusCode(v int32) *GetLatestSubmitDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLatestSubmitDetailResponse) SetBody(v *GetLatestSubmitDetailResponseBody) *GetLatestSubmitDetailResponse {
	s.Body = v
	return s
}

func (s *GetLatestSubmitDetailResponse) Validate() error {
	return dara.Validate(s)
}
