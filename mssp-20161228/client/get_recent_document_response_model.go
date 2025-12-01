// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecentDocumentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRecentDocumentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRecentDocumentResponse
	GetStatusCode() *int32
	SetBody(v *GetRecentDocumentResponseBody) *GetRecentDocumentResponse
	GetBody() *GetRecentDocumentResponseBody
}

type GetRecentDocumentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRecentDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRecentDocumentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRecentDocumentResponse) GoString() string {
	return s.String()
}

func (s *GetRecentDocumentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRecentDocumentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRecentDocumentResponse) GetBody() *GetRecentDocumentResponseBody {
	return s.Body
}

func (s *GetRecentDocumentResponse) SetHeaders(v map[string]*string) *GetRecentDocumentResponse {
	s.Headers = v
	return s
}

func (s *GetRecentDocumentResponse) SetStatusCode(v int32) *GetRecentDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecentDocumentResponse) SetBody(v *GetRecentDocumentResponseBody) *GetRecentDocumentResponse {
	s.Body = v
	return s
}

func (s *GetRecentDocumentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
