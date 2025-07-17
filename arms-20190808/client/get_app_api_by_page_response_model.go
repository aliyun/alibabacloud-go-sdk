// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppApiByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppApiByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppApiByPageResponse
	GetStatusCode() *int32
	SetBody(v *GetAppApiByPageResponseBody) *GetAppApiByPageResponse
	GetBody() *GetAppApiByPageResponseBody
}

type GetAppApiByPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppApiByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppApiByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppApiByPageResponse) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppApiByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppApiByPageResponse) GetBody() *GetAppApiByPageResponseBody {
	return s.Body
}

func (s *GetAppApiByPageResponse) SetHeaders(v map[string]*string) *GetAppApiByPageResponse {
	s.Headers = v
	return s
}

func (s *GetAppApiByPageResponse) SetStatusCode(v int32) *GetAppApiByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppApiByPageResponse) SetBody(v *GetAppApiByPageResponseBody) *GetAppApiByPageResponse {
	s.Body = v
	return s
}

func (s *GetAppApiByPageResponse) Validate() error {
	return dara.Validate(s)
}
