// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReindexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReindexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReindexResponse
	GetStatusCode() *int32
	SetBody(v *ReindexResponseBody) *ReindexResponse
	GetBody() *ReindexResponseBody
}

type ReindexResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReindexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReindexResponse) String() string {
	return dara.Prettify(s)
}

func (s ReindexResponse) GoString() string {
	return s.String()
}

func (s *ReindexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReindexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReindexResponse) GetBody() *ReindexResponseBody {
	return s.Body
}

func (s *ReindexResponse) SetHeaders(v map[string]*string) *ReindexResponse {
	s.Headers = v
	return s
}

func (s *ReindexResponse) SetStatusCode(v int32) *ReindexResponse {
	s.StatusCode = &v
	return s
}

func (s *ReindexResponse) SetBody(v *ReindexResponseBody) *ReindexResponse {
	s.Body = v
	return s
}

func (s *ReindexResponse) Validate() error {
	return dara.Validate(s)
}
