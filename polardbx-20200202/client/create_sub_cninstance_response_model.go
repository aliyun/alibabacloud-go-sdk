// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCNInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubCNInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubCNInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubCNInstanceResponseBody) *CreateSubCNInstanceResponse
	GetBody() *CreateSubCNInstanceResponseBody
}

type CreateSubCNInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubCNInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubCNInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCNInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSubCNInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubCNInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubCNInstanceResponse) GetBody() *CreateSubCNInstanceResponseBody {
	return s.Body
}

func (s *CreateSubCNInstanceResponse) SetHeaders(v map[string]*string) *CreateSubCNInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSubCNInstanceResponse) SetStatusCode(v int32) *CreateSubCNInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubCNInstanceResponse) SetBody(v *CreateSubCNInstanceResponseBody) *CreateSubCNInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateSubCNInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
