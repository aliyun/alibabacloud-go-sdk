// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSyntheticTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSyntheticTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSyntheticTaskListResponse
	GetStatusCode() *int32
	SetBody(v *GetSyntheticTaskListResponseBody) *GetSyntheticTaskListResponse
	GetBody() *GetSyntheticTaskListResponseBody
}

type GetSyntheticTaskListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSyntheticTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSyntheticTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSyntheticTaskListResponse) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSyntheticTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSyntheticTaskListResponse) GetBody() *GetSyntheticTaskListResponseBody {
	return s.Body
}

func (s *GetSyntheticTaskListResponse) SetHeaders(v map[string]*string) *GetSyntheticTaskListResponse {
	s.Headers = v
	return s
}

func (s *GetSyntheticTaskListResponse) SetStatusCode(v int32) *GetSyntheticTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyntheticTaskListResponse) SetBody(v *GetSyntheticTaskListResponseBody) *GetSyntheticTaskListResponse {
	s.Body = v
	return s
}

func (s *GetSyntheticTaskListResponse) Validate() error {
	return dara.Validate(s)
}
