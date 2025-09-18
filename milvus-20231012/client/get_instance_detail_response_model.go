// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceDetailResponseBody) *GetInstanceDetailResponse
	GetBody() *GetInstanceDetailResponseBody
}

type GetInstanceDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceDetailResponse) GetBody() *GetInstanceDetailResponseBody {
	return s.Body
}

func (s *GetInstanceDetailResponse) SetHeaders(v map[string]*string) *GetInstanceDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceDetailResponse) SetStatusCode(v int32) *GetInstanceDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceDetailResponse) SetBody(v *GetInstanceDetailResponseBody) *GetInstanceDetailResponse {
	s.Body = v
	return s
}

func (s *GetInstanceDetailResponse) Validate() error {
	return dara.Validate(s)
}
