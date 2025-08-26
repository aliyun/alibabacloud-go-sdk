// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAirflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAirflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAirflowResponse
	GetStatusCode() *int32
	SetBody(v *CreateAirflowResponseBody) *CreateAirflowResponse
	GetBody() *CreateAirflowResponseBody
}

type CreateAirflowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAirflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAirflowResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowResponse) GoString() string {
	return s.String()
}

func (s *CreateAirflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAirflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAirflowResponse) GetBody() *CreateAirflowResponseBody {
	return s.Body
}

func (s *CreateAirflowResponse) SetHeaders(v map[string]*string) *CreateAirflowResponse {
	s.Headers = v
	return s
}

func (s *CreateAirflowResponse) SetStatusCode(v int32) *CreateAirflowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAirflowResponse) SetBody(v *CreateAirflowResponseBody) *CreateAirflowResponse {
	s.Body = v
	return s
}

func (s *CreateAirflowResponse) Validate() error {
	return dara.Validate(s)
}
