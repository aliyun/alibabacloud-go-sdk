// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageDetailResponseBody) *GetMessageDetailResponse
	GetBody() *GetMessageDetailResponseBody
}

type GetMessageDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMessageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageDetailResponse) GetBody() *GetMessageDetailResponseBody {
	return s.Body
}

func (s *GetMessageDetailResponse) SetHeaders(v map[string]*string) *GetMessageDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMessageDetailResponse) SetStatusCode(v int32) *GetMessageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageDetailResponse) SetBody(v *GetMessageDetailResponseBody) *GetMessageDetailResponse {
	s.Body = v
	return s
}

func (s *GetMessageDetailResponse) Validate() error {
	return dara.Validate(s)
}
