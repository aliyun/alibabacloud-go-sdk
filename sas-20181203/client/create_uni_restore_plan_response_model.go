// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUniRestorePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUniRestorePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUniRestorePlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateUniRestorePlanResponseBody) *CreateUniRestorePlanResponse
	GetBody() *CreateUniRestorePlanResponseBody
}

type CreateUniRestorePlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUniRestorePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUniRestorePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUniRestorePlanResponse) GoString() string {
	return s.String()
}

func (s *CreateUniRestorePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUniRestorePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUniRestorePlanResponse) GetBody() *CreateUniRestorePlanResponseBody {
	return s.Body
}

func (s *CreateUniRestorePlanResponse) SetHeaders(v map[string]*string) *CreateUniRestorePlanResponse {
	s.Headers = v
	return s
}

func (s *CreateUniRestorePlanResponse) SetStatusCode(v int32) *CreateUniRestorePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUniRestorePlanResponse) SetBody(v *CreateUniRestorePlanResponseBody) *CreateUniRestorePlanResponse {
	s.Body = v
	return s
}

func (s *CreateUniRestorePlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
