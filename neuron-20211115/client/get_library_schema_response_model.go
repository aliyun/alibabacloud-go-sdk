// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibrarySchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLibrarySchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLibrarySchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetLibrarySchemaResponseBody) *GetLibrarySchemaResponse
	GetBody() *GetLibrarySchemaResponseBody
}

type GetLibrarySchemaResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibrarySchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibrarySchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLibrarySchemaResponse) GoString() string {
	return s.String()
}

func (s *GetLibrarySchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLibrarySchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLibrarySchemaResponse) GetBody() *GetLibrarySchemaResponseBody {
	return s.Body
}

func (s *GetLibrarySchemaResponse) SetHeaders(v map[string]*string) *GetLibrarySchemaResponse {
	s.Headers = v
	return s
}

func (s *GetLibrarySchemaResponse) SetStatusCode(v int32) *GetLibrarySchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibrarySchemaResponse) SetBody(v *GetLibrarySchemaResponseBody) *GetLibrarySchemaResponse {
	s.Body = v
	return s
}

func (s *GetLibrarySchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
