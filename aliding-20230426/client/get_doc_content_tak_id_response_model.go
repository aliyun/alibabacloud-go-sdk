// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentTakIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocContentTakIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocContentTakIdResponse
	GetStatusCode() *int32
	SetBody(v *GetDocContentTakIdResponseBody) *GetDocContentTakIdResponse
	GetBody() *GetDocContentTakIdResponseBody
}

type GetDocContentTakIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocContentTakIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocContentTakIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentTakIdResponse) GoString() string {
	return s.String()
}

func (s *GetDocContentTakIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocContentTakIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocContentTakIdResponse) GetBody() *GetDocContentTakIdResponseBody {
	return s.Body
}

func (s *GetDocContentTakIdResponse) SetHeaders(v map[string]*string) *GetDocContentTakIdResponse {
	s.Headers = v
	return s
}

func (s *GetDocContentTakIdResponse) SetStatusCode(v int32) *GetDocContentTakIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocContentTakIdResponse) SetBody(v *GetDocContentTakIdResponseBody) *GetDocContentTakIdResponse {
	s.Body = v
	return s
}

func (s *GetDocContentTakIdResponse) Validate() error {
	return dara.Validate(s)
}
