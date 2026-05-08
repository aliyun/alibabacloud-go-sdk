// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAICoachTaskSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseAICoachTaskSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseAICoachTaskSessionResponse
	GetStatusCode() *int32
	SetBody(v *CloseAICoachTaskSessionResponseBody) *CloseAICoachTaskSessionResponse
	GetBody() *CloseAICoachTaskSessionResponseBody
}

type CloseAICoachTaskSessionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseAICoachTaskSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseAICoachTaskSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseAICoachTaskSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseAICoachTaskSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseAICoachTaskSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseAICoachTaskSessionResponse) GetBody() *CloseAICoachTaskSessionResponseBody {
	return s.Body
}

func (s *CloseAICoachTaskSessionResponse) SetHeaders(v map[string]*string) *CloseAICoachTaskSessionResponse {
	s.Headers = v
	return s
}

func (s *CloseAICoachTaskSessionResponse) SetStatusCode(v int32) *CloseAICoachTaskSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseAICoachTaskSessionResponse) SetBody(v *CloseAICoachTaskSessionResponseBody) *CloseAICoachTaskSessionResponse {
	s.Body = v
	return s
}

func (s *CloseAICoachTaskSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
