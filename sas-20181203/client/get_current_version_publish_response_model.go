// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentVersionPublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCurrentVersionPublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCurrentVersionPublishResponse
	GetStatusCode() *int32
	SetBody(v *GetCurrentVersionPublishResponseBody) *GetCurrentVersionPublishResponse
	GetBody() *GetCurrentVersionPublishResponseBody
}

type GetCurrentVersionPublishResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCurrentVersionPublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCurrentVersionPublishResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentVersionPublishResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentVersionPublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCurrentVersionPublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCurrentVersionPublishResponse) GetBody() *GetCurrentVersionPublishResponseBody {
	return s.Body
}

func (s *GetCurrentVersionPublishResponse) SetHeaders(v map[string]*string) *GetCurrentVersionPublishResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentVersionPublishResponse) SetStatusCode(v int32) *GetCurrentVersionPublishResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCurrentVersionPublishResponse) SetBody(v *GetCurrentVersionPublishResponseBody) *GetCurrentVersionPublishResponse {
	s.Body = v
	return s
}

func (s *GetCurrentVersionPublishResponse) Validate() error {
	return dara.Validate(s)
}
