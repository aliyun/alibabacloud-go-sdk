// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentListResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentListResponseBody) *GetDocumentListResponse
	GetBody() *GetDocumentListResponseBody
}

type GetDocumentListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentListResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentListResponse) GetBody() *GetDocumentListResponseBody {
	return s.Body
}

func (s *GetDocumentListResponse) SetHeaders(v map[string]*string) *GetDocumentListResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentListResponse) SetStatusCode(v int32) *GetDocumentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentListResponse) SetBody(v *GetDocumentListResponseBody) *GetDocumentListResponse {
	s.Body = v
	return s
}

func (s *GetDocumentListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
