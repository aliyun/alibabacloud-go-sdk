// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportTaskNumberDatasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportTaskNumberDatasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportTaskNumberDatasResponse
	GetStatusCode() *int32
	SetBody(v *ImportTaskNumberDatasResponseBody) *ImportTaskNumberDatasResponse
	GetBody() *ImportTaskNumberDatasResponseBody
}

type ImportTaskNumberDatasResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportTaskNumberDatasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportTaskNumberDatasResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportTaskNumberDatasResponse) GoString() string {
	return s.String()
}

func (s *ImportTaskNumberDatasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportTaskNumberDatasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportTaskNumberDatasResponse) GetBody() *ImportTaskNumberDatasResponseBody {
	return s.Body
}

func (s *ImportTaskNumberDatasResponse) SetHeaders(v map[string]*string) *ImportTaskNumberDatasResponse {
	s.Headers = v
	return s
}

func (s *ImportTaskNumberDatasResponse) SetStatusCode(v int32) *ImportTaskNumberDatasResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportTaskNumberDatasResponse) SetBody(v *ImportTaskNumberDatasResponseBody) *ImportTaskNumberDatasResponse {
	s.Body = v
	return s
}

func (s *ImportTaskNumberDatasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
