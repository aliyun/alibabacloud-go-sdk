// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInventoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInventoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInventoryJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateInventoryJobResponseBody) *CreateInventoryJobResponse
	GetBody() *CreateInventoryJobResponseBody
}

type CreateInventoryJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInventoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInventoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInventoryJobResponse) GoString() string {
	return s.String()
}

func (s *CreateInventoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInventoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInventoryJobResponse) GetBody() *CreateInventoryJobResponseBody {
	return s.Body
}

func (s *CreateInventoryJobResponse) SetHeaders(v map[string]*string) *CreateInventoryJobResponse {
	s.Headers = v
	return s
}

func (s *CreateInventoryJobResponse) SetStatusCode(v int32) *CreateInventoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInventoryJobResponse) SetBody(v *CreateInventoryJobResponseBody) *CreateInventoryJobResponse {
	s.Body = v
	return s
}

func (s *CreateInventoryJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
