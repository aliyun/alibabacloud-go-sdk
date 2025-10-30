// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDocumentCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDocumentCollectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateDocumentCollectionResponseBody) *CreateDocumentCollectionResponse
	GetBody() *CreateDocumentCollectionResponseBody
}

type CreateDocumentCollectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDocumentCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDocumentCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentCollectionResponse) GoString() string {
	return s.String()
}

func (s *CreateDocumentCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDocumentCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDocumentCollectionResponse) GetBody() *CreateDocumentCollectionResponseBody {
	return s.Body
}

func (s *CreateDocumentCollectionResponse) SetHeaders(v map[string]*string) *CreateDocumentCollectionResponse {
	s.Headers = v
	return s
}

func (s *CreateDocumentCollectionResponse) SetStatusCode(v int32) *CreateDocumentCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDocumentCollectionResponse) SetBody(v *CreateDocumentCollectionResponseBody) *CreateDocumentCollectionResponse {
	s.Body = v
	return s
}

func (s *CreateDocumentCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
