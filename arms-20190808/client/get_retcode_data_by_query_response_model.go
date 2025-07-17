// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeDataByQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRetcodeDataByQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRetcodeDataByQueryResponse
	GetStatusCode() *int32
	SetBody(v *GetRetcodeDataByQueryResponseBody) *GetRetcodeDataByQueryResponse
	GetBody() *GetRetcodeDataByQueryResponseBody
}

type GetRetcodeDataByQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRetcodeDataByQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRetcodeDataByQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeDataByQueryResponse) GoString() string {
	return s.String()
}

func (s *GetRetcodeDataByQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRetcodeDataByQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRetcodeDataByQueryResponse) GetBody() *GetRetcodeDataByQueryResponseBody {
	return s.Body
}

func (s *GetRetcodeDataByQueryResponse) SetHeaders(v map[string]*string) *GetRetcodeDataByQueryResponse {
	s.Headers = v
	return s
}

func (s *GetRetcodeDataByQueryResponse) SetStatusCode(v int32) *GetRetcodeDataByQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRetcodeDataByQueryResponse) SetBody(v *GetRetcodeDataByQueryResponseBody) *GetRetcodeDataByQueryResponse {
	s.Body = v
	return s
}

func (s *GetRetcodeDataByQueryResponse) Validate() error {
	return dara.Validate(s)
}
