// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibrarySchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLibrarySchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLibrarySchemaResponse
	GetStatusCode() *int32
	SetBody(v *LibrarySchema) *CreateLibrarySchemaResponse
	GetBody() *LibrarySchema
}

type CreateLibrarySchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibrarySchema     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibrarySchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLibrarySchemaResponse) GoString() string {
	return s.String()
}

func (s *CreateLibrarySchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLibrarySchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLibrarySchemaResponse) GetBody() *LibrarySchema {
	return s.Body
}

func (s *CreateLibrarySchemaResponse) SetHeaders(v map[string]*string) *CreateLibrarySchemaResponse {
	s.Headers = v
	return s
}

func (s *CreateLibrarySchemaResponse) SetStatusCode(v int32) *CreateLibrarySchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLibrarySchemaResponse) SetBody(v *LibrarySchema) *CreateLibrarySchemaResponse {
	s.Body = v
	return s
}

func (s *CreateLibrarySchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
