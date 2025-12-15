// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentSplitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocumentSplitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocumentSplitResponse
	GetStatusCode() *int32
	SetBody(v *GetDocumentSplitResponseBody) *GetDocumentSplitResponse
	GetBody() *GetDocumentSplitResponseBody
}

type GetDocumentSplitResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentSplitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentSplitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocumentSplitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocumentSplitResponse) GetBody() *GetDocumentSplitResponseBody {
	return s.Body
}

func (s *GetDocumentSplitResponse) SetHeaders(v map[string]*string) *GetDocumentSplitResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentSplitResponse) SetStatusCode(v int32) *GetDocumentSplitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentSplitResponse) SetBody(v *GetDocumentSplitResponseBody) *GetDocumentSplitResponse {
	s.Body = v
	return s
}

func (s *GetDocumentSplitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
