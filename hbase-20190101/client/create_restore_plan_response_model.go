// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestorePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRestorePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRestorePlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateRestorePlanResponseBody) *CreateRestorePlanResponse
	GetBody() *CreateRestorePlanResponseBody
}

type CreateRestorePlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRestorePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRestorePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRestorePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateRestorePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRestorePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRestorePlanResponse) GetBody() *CreateRestorePlanResponseBody {
	return s.Body
}

func (s *CreateRestorePlanResponse) SetHeaders(v map[string]*string) *CreateRestorePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateRestorePlanResponse) SetStatusCode(v int32) *CreateRestorePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRestorePlanResponse) SetBody(v *CreateRestorePlanResponseBody) *CreateRestorePlanResponse {
	s.Body = v
	return s
}

func (s *CreateRestorePlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
