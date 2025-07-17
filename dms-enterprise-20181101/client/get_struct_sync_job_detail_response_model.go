// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStructSyncJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStructSyncJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetStructSyncJobDetailResponseBody) *GetStructSyncJobDetailResponse
	GetBody() *GetStructSyncJobDetailResponseBody
}

type GetStructSyncJobDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStructSyncJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStructSyncJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStructSyncJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStructSyncJobDetailResponse) GetBody() *GetStructSyncJobDetailResponseBody {
	return s.Body
}

func (s *GetStructSyncJobDetailResponse) SetHeaders(v map[string]*string) *GetStructSyncJobDetailResponse {
	s.Headers = v
	return s
}

func (s *GetStructSyncJobDetailResponse) SetStatusCode(v int32) *GetStructSyncJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStructSyncJobDetailResponse) SetBody(v *GetStructSyncJobDetailResponseBody) *GetStructSyncJobDetailResponse {
	s.Body = v
	return s
}

func (s *GetStructSyncJobDetailResponse) Validate() error {
	return dara.Validate(s)
}
