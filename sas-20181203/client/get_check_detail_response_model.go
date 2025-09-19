// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckDetailResponseBody) *GetCheckDetailResponse
	GetBody() *GetCheckDetailResponseBody
}

type GetCheckDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCheckDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckDetailResponse) GetBody() *GetCheckDetailResponseBody {
	return s.Body
}

func (s *GetCheckDetailResponse) SetHeaders(v map[string]*string) *GetCheckDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCheckDetailResponse) SetStatusCode(v int32) *GetCheckDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckDetailResponse) SetBody(v *GetCheckDetailResponseBody) *GetCheckDetailResponse {
	s.Body = v
	return s
}

func (s *GetCheckDetailResponse) Validate() error {
	return dara.Validate(s)
}
