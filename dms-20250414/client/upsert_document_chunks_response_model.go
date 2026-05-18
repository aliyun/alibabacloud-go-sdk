// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertDocumentChunksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpsertDocumentChunksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpsertDocumentChunksResponse
	GetStatusCode() *int32
	SetBody(v *UpsertDocumentChunksResponseBody) *UpsertDocumentChunksResponse
	GetBody() *UpsertDocumentChunksResponseBody
}

type UpsertDocumentChunksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpsertDocumentChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpsertDocumentChunksResponse) String() string {
	return dara.Prettify(s)
}

func (s UpsertDocumentChunksResponse) GoString() string {
	return s.String()
}

func (s *UpsertDocumentChunksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpsertDocumentChunksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpsertDocumentChunksResponse) GetBody() *UpsertDocumentChunksResponseBody {
	return s.Body
}

func (s *UpsertDocumentChunksResponse) SetHeaders(v map[string]*string) *UpsertDocumentChunksResponse {
	s.Headers = v
	return s
}

func (s *UpsertDocumentChunksResponse) SetStatusCode(v int32) *UpsertDocumentChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *UpsertDocumentChunksResponse) SetBody(v *UpsertDocumentChunksResponseBody) *UpsertDocumentChunksResponse {
	s.Body = v
	return s
}

func (s *UpsertDocumentChunksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
