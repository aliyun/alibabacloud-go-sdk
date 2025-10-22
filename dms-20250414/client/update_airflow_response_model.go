// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAirflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAirflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAirflowResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAirflowResponseBody) *UpdateAirflowResponse
	GetBody() *UpdateAirflowResponseBody
}

type UpdateAirflowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAirflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAirflowResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAirflowResponse) GoString() string {
	return s.String()
}

func (s *UpdateAirflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAirflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAirflowResponse) GetBody() *UpdateAirflowResponseBody {
	return s.Body
}

func (s *UpdateAirflowResponse) SetHeaders(v map[string]*string) *UpdateAirflowResponse {
	s.Headers = v
	return s
}

func (s *UpdateAirflowResponse) SetStatusCode(v int32) *UpdateAirflowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAirflowResponse) SetBody(v *UpdateAirflowResponseBody) *UpdateAirflowResponse {
	s.Body = v
	return s
}

func (s *UpdateAirflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
