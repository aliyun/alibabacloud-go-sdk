// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCrowdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubCrowdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubCrowdResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubCrowdResponseBody) *CreateSubCrowdResponse
	GetBody() *CreateSubCrowdResponseBody
}

type CreateSubCrowdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubCrowdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubCrowdResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCrowdResponse) GoString() string {
	return s.String()
}

func (s *CreateSubCrowdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubCrowdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubCrowdResponse) GetBody() *CreateSubCrowdResponseBody {
	return s.Body
}

func (s *CreateSubCrowdResponse) SetHeaders(v map[string]*string) *CreateSubCrowdResponse {
	s.Headers = v
	return s
}

func (s *CreateSubCrowdResponse) SetStatusCode(v int32) *CreateSubCrowdResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubCrowdResponse) SetBody(v *CreateSubCrowdResponseBody) *CreateSubCrowdResponse {
	s.Body = v
	return s
}

func (s *CreateSubCrowdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
