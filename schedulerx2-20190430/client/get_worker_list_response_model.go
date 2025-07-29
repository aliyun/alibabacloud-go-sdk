// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkerListResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkerListResponseBody) *GetWorkerListResponse
	GetBody() *GetWorkerListResponseBody
}

type GetWorkerListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkerListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerListResponse) GoString() string {
	return s.String()
}

func (s *GetWorkerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkerListResponse) GetBody() *GetWorkerListResponseBody {
	return s.Body
}

func (s *GetWorkerListResponse) SetHeaders(v map[string]*string) *GetWorkerListResponse {
	s.Headers = v
	return s
}

func (s *GetWorkerListResponse) SetStatusCode(v int32) *GetWorkerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkerListResponse) SetBody(v *GetWorkerListResponseBody) *GetWorkerListResponse {
	s.Body = v
	return s
}

func (s *GetWorkerListResponse) Validate() error {
	return dara.Validate(s)
}
