// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLibrarySchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLibrarySchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLibrarySchemaResponse
	GetStatusCode() *int32
}

type DeleteLibrarySchemaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteLibrarySchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLibrarySchemaResponse) GoString() string {
	return s.String()
}

func (s *DeleteLibrarySchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLibrarySchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLibrarySchemaResponse) SetHeaders(v map[string]*string) *DeleteLibrarySchemaResponse {
	s.Headers = v
	return s
}

func (s *DeleteLibrarySchemaResponse) SetStatusCode(v int32) *DeleteLibrarySchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLibrarySchemaResponse) Validate() error {
	return dara.Validate(s)
}
