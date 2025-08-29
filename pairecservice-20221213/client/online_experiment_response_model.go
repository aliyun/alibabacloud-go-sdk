// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineExperimentResponse
	GetStatusCode() *int32
	SetBody(v *OnlineExperimentResponseBody) *OnlineExperimentResponse
	GetBody() *OnlineExperimentResponseBody
}

type OnlineExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineExperimentResponse) GoString() string {
	return s.String()
}

func (s *OnlineExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineExperimentResponse) GetBody() *OnlineExperimentResponseBody {
	return s.Body
}

func (s *OnlineExperimentResponse) SetHeaders(v map[string]*string) *OnlineExperimentResponse {
	s.Headers = v
	return s
}

func (s *OnlineExperimentResponse) SetStatusCode(v int32) *OnlineExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineExperimentResponse) SetBody(v *OnlineExperimentResponseBody) *OnlineExperimentResponse {
	s.Body = v
	return s
}

func (s *OnlineExperimentResponse) Validate() error {
	return dara.Validate(s)
}
