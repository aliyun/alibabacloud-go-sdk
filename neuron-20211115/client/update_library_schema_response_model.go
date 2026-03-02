// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibrarySchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLibrarySchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLibrarySchemaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLibrarySchemaResponseBody) *UpdateLibrarySchemaResponse
	GetBody() *UpdateLibrarySchemaResponseBody
}

type UpdateLibrarySchemaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLibrarySchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibrarySchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLibrarySchemaResponse) GoString() string {
	return s.String()
}

func (s *UpdateLibrarySchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLibrarySchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLibrarySchemaResponse) GetBody() *UpdateLibrarySchemaResponseBody {
	return s.Body
}

func (s *UpdateLibrarySchemaResponse) SetHeaders(v map[string]*string) *UpdateLibrarySchemaResponse {
	s.Headers = v
	return s
}

func (s *UpdateLibrarySchemaResponse) SetStatusCode(v int32) *UpdateLibrarySchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLibrarySchemaResponse) SetBody(v *UpdateLibrarySchemaResponseBody) *UpdateLibrarySchemaResponse {
	s.Body = v
	return s
}

func (s *UpdateLibrarySchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
