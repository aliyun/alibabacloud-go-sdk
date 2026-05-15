// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployAirflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedeployAirflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedeployAirflowResponse
	GetStatusCode() *int32
	SetBody(v *RedeployAirflowResponseBody) *RedeployAirflowResponse
	GetBody() *RedeployAirflowResponseBody
}

type RedeployAirflowResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedeployAirflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedeployAirflowResponse) String() string {
	return dara.Prettify(s)
}

func (s RedeployAirflowResponse) GoString() string {
	return s.String()
}

func (s *RedeployAirflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedeployAirflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedeployAirflowResponse) GetBody() *RedeployAirflowResponseBody {
	return s.Body
}

func (s *RedeployAirflowResponse) SetHeaders(v map[string]*string) *RedeployAirflowResponse {
	s.Headers = v
	return s
}

func (s *RedeployAirflowResponse) SetStatusCode(v int32) *RedeployAirflowResponse {
	s.StatusCode = &v
	return s
}

func (s *RedeployAirflowResponse) SetBody(v *RedeployAirflowResponseBody) *RedeployAirflowResponse {
	s.Body = v
	return s
}

func (s *RedeployAirflowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
