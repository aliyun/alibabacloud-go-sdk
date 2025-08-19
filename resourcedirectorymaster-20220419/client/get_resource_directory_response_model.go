// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceDirectoryResponseBody) *GetResourceDirectoryResponse
	GetBody() *GetResourceDirectoryResponseBody
}

type GetResourceDirectoryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceDirectoryResponse) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceDirectoryResponse) GetBody() *GetResourceDirectoryResponseBody {
	return s.Body
}

func (s *GetResourceDirectoryResponse) SetHeaders(v map[string]*string) *GetResourceDirectoryResponse {
	s.Headers = v
	return s
}

func (s *GetResourceDirectoryResponse) SetStatusCode(v int32) *GetResourceDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceDirectoryResponse) SetBody(v *GetResourceDirectoryResponseBody) *GetResourceDirectoryResponse {
	s.Body = v
	return s
}

func (s *GetResourceDirectoryResponse) Validate() error {
	return dara.Validate(s)
}
