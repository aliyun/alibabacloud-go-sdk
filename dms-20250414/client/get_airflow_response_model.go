// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAirflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAirflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAirflowResponse
	GetStatusCode() *int32
	SetBody(v *GetAirflowResponseBody) *GetAirflowResponse
	GetBody() *GetAirflowResponseBody
}

type GetAirflowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAirflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAirflowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAirflowResponse) GoString() string {
	return s.String()
}

func (s *GetAirflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAirflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAirflowResponse) GetBody() *GetAirflowResponseBody {
	return s.Body
}

func (s *GetAirflowResponse) SetHeaders(v map[string]*string) *GetAirflowResponse {
	s.Headers = v
	return s
}

func (s *GetAirflowResponse) SetStatusCode(v int32) *GetAirflowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAirflowResponse) SetBody(v *GetAirflowResponseBody) *GetAirflowResponse {
	s.Body = v
	return s
}

func (s *GetAirflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
