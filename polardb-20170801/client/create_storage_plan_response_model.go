// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoragePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStoragePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStoragePlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateStoragePlanResponseBody) *CreateStoragePlanResponse
	GetBody() *CreateStoragePlanResponseBody
}

type CreateStoragePlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStoragePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStoragePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStoragePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateStoragePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStoragePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStoragePlanResponse) GetBody() *CreateStoragePlanResponseBody {
	return s.Body
}

func (s *CreateStoragePlanResponse) SetHeaders(v map[string]*string) *CreateStoragePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateStoragePlanResponse) SetStatusCode(v int32) *CreateStoragePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoragePlanResponse) SetBody(v *CreateStoragePlanResponseBody) *CreateStoragePlanResponse {
	s.Body = v
	return s
}

func (s *CreateStoragePlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
