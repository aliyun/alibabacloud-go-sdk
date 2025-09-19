// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetLogMetaResponseBody) *GetLogMetaResponse
	GetBody() *GetLogMetaResponseBody
}

type GetLogMetaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogMetaResponse) GoString() string {
	return s.String()
}

func (s *GetLogMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogMetaResponse) GetBody() *GetLogMetaResponseBody {
	return s.Body
}

func (s *GetLogMetaResponse) SetHeaders(v map[string]*string) *GetLogMetaResponse {
	s.Headers = v
	return s
}

func (s *GetLogMetaResponse) SetStatusCode(v int32) *GetLogMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogMetaResponse) SetBody(v *GetLogMetaResponseBody) *GetLogMetaResponse {
	s.Body = v
	return s
}

func (s *GetLogMetaResponse) Validate() error {
	return dara.Validate(s)
}
