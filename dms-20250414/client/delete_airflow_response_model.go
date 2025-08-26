// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAirflowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAirflowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAirflowResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAirflowResponseBody) *DeleteAirflowResponse
	GetBody() *DeleteAirflowResponseBody
}

type DeleteAirflowResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAirflowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAirflowResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAirflowResponse) GoString() string {
	return s.String()
}

func (s *DeleteAirflowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAirflowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAirflowResponse) GetBody() *DeleteAirflowResponseBody {
	return s.Body
}

func (s *DeleteAirflowResponse) SetHeaders(v map[string]*string) *DeleteAirflowResponse {
	s.Headers = v
	return s
}

func (s *DeleteAirflowResponse) SetStatusCode(v int32) *DeleteAirflowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAirflowResponse) SetBody(v *DeleteAirflowResponseBody) *DeleteAirflowResponse {
	s.Body = v
	return s
}

func (s *DeleteAirflowResponse) Validate() error {
	return dara.Validate(s)
}
