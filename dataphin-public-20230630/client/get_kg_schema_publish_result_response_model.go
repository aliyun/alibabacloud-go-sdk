// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgSchemaPublishResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKgSchemaPublishResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKgSchemaPublishResultResponse
	GetStatusCode() *int32
	SetBody(v *GetKgSchemaPublishResultResponseBody) *GetKgSchemaPublishResultResponse
	GetBody() *GetKgSchemaPublishResultResponseBody
}

type GetKgSchemaPublishResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKgSchemaPublishResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKgSchemaPublishResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKgSchemaPublishResultResponse) GoString() string {
	return s.String()
}

func (s *GetKgSchemaPublishResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKgSchemaPublishResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKgSchemaPublishResultResponse) GetBody() *GetKgSchemaPublishResultResponseBody {
	return s.Body
}

func (s *GetKgSchemaPublishResultResponse) SetHeaders(v map[string]*string) *GetKgSchemaPublishResultResponse {
	s.Headers = v
	return s
}

func (s *GetKgSchemaPublishResultResponse) SetStatusCode(v int32) *GetKgSchemaPublishResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKgSchemaPublishResultResponse) SetBody(v *GetKgSchemaPublishResultResponseBody) *GetKgSchemaPublishResultResponse {
	s.Body = v
	return s
}

func (s *GetKgSchemaPublishResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
