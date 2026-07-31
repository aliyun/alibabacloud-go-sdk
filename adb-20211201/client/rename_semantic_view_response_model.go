// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameSemanticViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenameSemanticViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenameSemanticViewResponse
	GetStatusCode() *int32
	SetBody(v *RenameSemanticViewResponseBody) *RenameSemanticViewResponse
	GetBody() *RenameSemanticViewResponseBody
}

type RenameSemanticViewResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameSemanticViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameSemanticViewResponse) String() string {
	return dara.Prettify(s)
}

func (s RenameSemanticViewResponse) GoString() string {
	return s.String()
}

func (s *RenameSemanticViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenameSemanticViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenameSemanticViewResponse) GetBody() *RenameSemanticViewResponseBody {
	return s.Body
}

func (s *RenameSemanticViewResponse) SetHeaders(v map[string]*string) *RenameSemanticViewResponse {
	s.Headers = v
	return s
}

func (s *RenameSemanticViewResponse) SetStatusCode(v int32) *RenameSemanticViewResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameSemanticViewResponse) SetBody(v *RenameSemanticViewResponseBody) *RenameSemanticViewResponse {
	s.Body = v
	return s
}

func (s *RenameSemanticViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
