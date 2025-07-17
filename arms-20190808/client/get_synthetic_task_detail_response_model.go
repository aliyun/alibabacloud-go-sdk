// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSyntheticTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSyntheticTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetSyntheticTaskDetailResponseBody) *GetSyntheticTaskDetailResponse
	GetBody() *GetSyntheticTaskDetailResponseBody
}

type GetSyntheticTaskDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSyntheticTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSyntheticTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSyntheticTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSyntheticTaskDetailResponse) GetBody() *GetSyntheticTaskDetailResponseBody {
	return s.Body
}

func (s *GetSyntheticTaskDetailResponse) SetHeaders(v map[string]*string) *GetSyntheticTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetSyntheticTaskDetailResponse) SetStatusCode(v int32) *GetSyntheticTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyntheticTaskDetailResponse) SetBody(v *GetSyntheticTaskDetailResponseBody) *GetSyntheticTaskDetailResponse {
	s.Body = v
	return s
}

func (s *GetSyntheticTaskDetailResponse) Validate() error {
	return dara.Validate(s)
}
