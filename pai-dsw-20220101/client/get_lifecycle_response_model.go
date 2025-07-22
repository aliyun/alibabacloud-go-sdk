// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLifecycleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLifecycleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLifecycleResponse
	GetStatusCode() *int32
	SetBody(v *GetLifecycleResponseBody) *GetLifecycleResponse
	GetBody() *GetLifecycleResponseBody
}

type GetLifecycleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLifecycleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLifecycleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLifecycleResponse) GoString() string {
	return s.String()
}

func (s *GetLifecycleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLifecycleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLifecycleResponse) GetBody() *GetLifecycleResponseBody {
	return s.Body
}

func (s *GetLifecycleResponse) SetHeaders(v map[string]*string) *GetLifecycleResponse {
	s.Headers = v
	return s
}

func (s *GetLifecycleResponse) SetStatusCode(v int32) *GetLifecycleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLifecycleResponse) SetBody(v *GetLifecycleResponseBody) *GetLifecycleResponse {
	s.Body = v
	return s
}

func (s *GetLifecycleResponse) Validate() error {
	return dara.Validate(s)
}
