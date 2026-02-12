// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTrendTopicInputTpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTrendTopicInputTpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTrendTopicInputTpsResponse
	GetStatusCode() *int32
	SetBody(v *OnsTrendTopicInputTpsResponseBody) *OnsTrendTopicInputTpsResponse
	GetBody() *OnsTrendTopicInputTpsResponseBody
}

type OnsTrendTopicInputTpsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTrendTopicInputTpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTrendTopicInputTpsResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTrendTopicInputTpsResponse) GoString() string {
	return s.String()
}

func (s *OnsTrendTopicInputTpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTrendTopicInputTpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTrendTopicInputTpsResponse) GetBody() *OnsTrendTopicInputTpsResponseBody {
	return s.Body
}

func (s *OnsTrendTopicInputTpsResponse) SetHeaders(v map[string]*string) *OnsTrendTopicInputTpsResponse {
	s.Headers = v
	return s
}

func (s *OnsTrendTopicInputTpsResponse) SetStatusCode(v int32) *OnsTrendTopicInputTpsResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTrendTopicInputTpsResponse) SetBody(v *OnsTrendTopicInputTpsResponseBody) *OnsTrendTopicInputTpsResponse {
	s.Body = v
	return s
}

func (s *OnsTrendTopicInputTpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
