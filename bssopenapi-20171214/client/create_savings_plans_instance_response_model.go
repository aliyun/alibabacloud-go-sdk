// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlansInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSavingsPlansInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSavingsPlansInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSavingsPlansInstanceResponseBody) *CreateSavingsPlansInstanceResponse
	GetBody() *CreateSavingsPlansInstanceResponseBody
}

type CreateSavingsPlansInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSavingsPlansInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSavingsPlansInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlansInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlansInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSavingsPlansInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSavingsPlansInstanceResponse) GetBody() *CreateSavingsPlansInstanceResponseBody {
	return s.Body
}

func (s *CreateSavingsPlansInstanceResponse) SetHeaders(v map[string]*string) *CreateSavingsPlansInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSavingsPlansInstanceResponse) SetStatusCode(v int32) *CreateSavingsPlansInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSavingsPlansInstanceResponse) SetBody(v *CreateSavingsPlansInstanceResponseBody) *CreateSavingsPlansInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateSavingsPlansInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
