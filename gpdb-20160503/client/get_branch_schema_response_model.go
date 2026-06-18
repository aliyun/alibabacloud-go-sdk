// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBranchSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBranchSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBranchSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetBranchSchemaResponseBody) *GetBranchSchemaResponse
	GetBody() *GetBranchSchemaResponseBody
}

type GetBranchSchemaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBranchSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBranchSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBranchSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetBranchSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBranchSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBranchSchemaResponse) GetBody() *GetBranchSchemaResponseBody {
	return s.Body
}

func (s *GetBranchSchemaResponse) SetHeaders(v map[string]*string) *GetBranchSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetBranchSchemaResponse) SetStatusCode(v int32) *GetBranchSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBranchSchemaResponse) SetBody(v *GetBranchSchemaResponseBody) *GetBranchSchemaResponse {
	s.Body = v
	return s
}

func (s *GetBranchSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
