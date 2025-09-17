// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMeteringDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushMeteringDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushMeteringDataResponse
	GetStatusCode() *int32
	SetBody(v *PushMeteringDataResponseBody) *PushMeteringDataResponse
	GetBody() *PushMeteringDataResponseBody
}

type PushMeteringDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMeteringDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMeteringDataResponse) String() string {
	return dara.Prettify(s)
}

func (s PushMeteringDataResponse) GoString() string {
	return s.String()
}

func (s *PushMeteringDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushMeteringDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushMeteringDataResponse) GetBody() *PushMeteringDataResponseBody {
	return s.Body
}

func (s *PushMeteringDataResponse) SetHeaders(v map[string]*string) *PushMeteringDataResponse {
	s.Headers = v
	return s
}

func (s *PushMeteringDataResponse) SetStatusCode(v int32) *PushMeteringDataResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMeteringDataResponse) SetBody(v *PushMeteringDataResponseBody) *PushMeteringDataResponse {
	s.Body = v
	return s
}

func (s *PushMeteringDataResponse) Validate() error {
	return dara.Validate(s)
}
