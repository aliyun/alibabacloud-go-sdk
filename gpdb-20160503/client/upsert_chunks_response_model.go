// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertChunksResponse
	GetStatusCode() *int32
	SetBody(v *UpsertChunksResponseBody) *UpsertChunksResponse
	GetBody() *UpsertChunksResponseBody
}

type UpsertChunksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertChunksResponse) GoString() string {
	return s.String()
}

func (s *UpsertChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertChunksResponse) GetBody() *UpsertChunksResponseBody {
	return s.Body
}

func (s *UpsertChunksResponse) SetHeaders(v map[string]*string) *UpsertChunksResponse {
	s.Headers = v
	return s
}

func (s *UpsertChunksResponse) SetStatusCode(v int32) *UpsertChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertChunksResponse) SetBody(v *UpsertChunksResponseBody) *UpsertChunksResponse {
	s.Body = v
	return s
}

func (s *UpsertChunksResponse) Validate() error {
	return dara.Validate(s)
}
