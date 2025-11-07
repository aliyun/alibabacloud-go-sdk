// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventorySchemaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInventorySchemaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInventorySchemaResponse
	GetStatusCode() *int32
	SetBody(v *GetInventorySchemaResponseBody) *GetInventorySchemaResponse
	GetBody() *GetInventorySchemaResponseBody
}

type GetInventorySchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInventorySchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInventorySchemaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySchemaResponse) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInventorySchemaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInventorySchemaResponse) GetBody() *GetInventorySchemaResponseBody {
	return s.Body
}

func (s *GetInventorySchemaResponse) SetHeaders(v map[string]*string) *GetInventorySchemaResponse {
	s.Headers = v
	return s
}

func (s *GetInventorySchemaResponse) SetStatusCode(v int32) *GetInventorySchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInventorySchemaResponse) SetBody(v *GetInventorySchemaResponseBody) *GetInventorySchemaResponse {
	s.Body = v
	return s
}

func (s *GetInventorySchemaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
