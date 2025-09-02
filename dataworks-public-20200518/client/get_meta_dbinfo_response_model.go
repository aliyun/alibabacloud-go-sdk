// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaDBInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaDBInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaDBInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaDBInfoResponseBody) *GetMetaDBInfoResponse
	GetBody() *GetMetaDBInfoResponseBody
}

type GetMetaDBInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaDBInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaDBInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMetaDBInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaDBInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaDBInfoResponse) GetBody() *GetMetaDBInfoResponseBody {
	return s.Body
}

func (s *GetMetaDBInfoResponse) SetHeaders(v map[string]*string) *GetMetaDBInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMetaDBInfoResponse) SetStatusCode(v int32) *GetMetaDBInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaDBInfoResponse) SetBody(v *GetMetaDBInfoResponseBody) *GetMetaDBInfoResponse {
	s.Body = v
	return s
}

func (s *GetMetaDBInfoResponse) Validate() error {
	return dara.Validate(s)
}
