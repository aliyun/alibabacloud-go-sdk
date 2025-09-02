// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableOutputResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableOutputResponseBody) *GetMetaTableOutputResponse
	GetBody() *GetMetaTableOutputResponseBody
}

type GetMetaTableOutputResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableOutputResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableOutputResponse) GetBody() *GetMetaTableOutputResponseBody {
	return s.Body
}

func (s *GetMetaTableOutputResponse) SetHeaders(v map[string]*string) *GetMetaTableOutputResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableOutputResponse) SetStatusCode(v int32) *GetMetaTableOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableOutputResponse) SetBody(v *GetMetaTableOutputResponseBody) *GetMetaTableOutputResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableOutputResponse) Validate() error {
	return dara.Validate(s)
}
