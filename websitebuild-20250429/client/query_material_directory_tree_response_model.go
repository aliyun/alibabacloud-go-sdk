// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMaterialDirectoryTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMaterialDirectoryTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMaterialDirectoryTreeResponse
	GetStatusCode() *int32
	SetBody(v *QueryMaterialDirectoryTreeResponseBody) *QueryMaterialDirectoryTreeResponse
	GetBody() *QueryMaterialDirectoryTreeResponseBody
}

type QueryMaterialDirectoryTreeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMaterialDirectoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMaterialDirectoryTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMaterialDirectoryTreeResponse) GoString() string {
	return s.String()
}

func (s *QueryMaterialDirectoryTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMaterialDirectoryTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMaterialDirectoryTreeResponse) GetBody() *QueryMaterialDirectoryTreeResponseBody {
	return s.Body
}

func (s *QueryMaterialDirectoryTreeResponse) SetHeaders(v map[string]*string) *QueryMaterialDirectoryTreeResponse {
	s.Headers = v
	return s
}

func (s *QueryMaterialDirectoryTreeResponse) SetStatusCode(v int32) *QueryMaterialDirectoryTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMaterialDirectoryTreeResponse) SetBody(v *QueryMaterialDirectoryTreeResponseBody) *QueryMaterialDirectoryTreeResponse {
	s.Body = v
	return s
}

func (s *QueryMaterialDirectoryTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
