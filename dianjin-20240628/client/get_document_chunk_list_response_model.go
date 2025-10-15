// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentChunkListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentChunkListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentChunkListResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentChunkListResponseBody) *GetDocumentChunkListResponse
	GetBody() *GetDocumentChunkListResponseBody
}

type GetDocumentChunkListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentChunkListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentChunkListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentChunkListResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentChunkListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentChunkListResponse) GetBody() *GetDocumentChunkListResponseBody {
	return s.Body
}

func (s *GetDocumentChunkListResponse) SetHeaders(v map[string]*string) *GetDocumentChunkListResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentChunkListResponse) SetStatusCode(v int32) *GetDocumentChunkListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentChunkListResponse) SetBody(v *GetDocumentChunkListResponseBody) *GetDocumentChunkListResponse {
	s.Body = v
	return s
}

func (s *GetDocumentChunkListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
