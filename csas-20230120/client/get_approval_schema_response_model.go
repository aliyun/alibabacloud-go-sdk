// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApprovalSchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApprovalSchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApprovalSchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetApprovalSchemaResponseBody) *GetApprovalSchemaResponse
	GetBody() *GetApprovalSchemaResponseBody
}

type GetApprovalSchemaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApprovalSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApprovalSchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApprovalSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetApprovalSchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApprovalSchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApprovalSchemaResponse) GetBody() *GetApprovalSchemaResponseBody {
	return s.Body
}

func (s *GetApprovalSchemaResponse) SetHeaders(v map[string]*string) *GetApprovalSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetApprovalSchemaResponse) SetStatusCode(v int32) *GetApprovalSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApprovalSchemaResponse) SetBody(v *GetApprovalSchemaResponseBody) *GetApprovalSchemaResponse {
	s.Body = v
	return s
}

func (s *GetApprovalSchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
