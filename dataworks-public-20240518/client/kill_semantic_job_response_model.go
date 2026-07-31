// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSemanticJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *KillSemanticJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *KillSemanticJobResponse
	GetStatusCode() *int32
	SetBody(v *KillSemanticJobResponseBody) *KillSemanticJobResponse
	GetBody() *KillSemanticJobResponseBody
}

type KillSemanticJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSemanticJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s KillSemanticJobResponse) String() string {
	return dara.Prettify(s)
}

func (s KillSemanticJobResponse) GoString() string {
	return s.String()
}

func (s *KillSemanticJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *KillSemanticJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *KillSemanticJobResponse) GetBody() *KillSemanticJobResponseBody {
	return s.Body
}

func (s *KillSemanticJobResponse) SetHeaders(v map[string]*string) *KillSemanticJobResponse {
	s.Headers = v
	return s
}

func (s *KillSemanticJobResponse) SetStatusCode(v int32) *KillSemanticJobResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSemanticJobResponse) SetBody(v *KillSemanticJobResponseBody) *KillSemanticJobResponse {
	s.Body = v
	return s
}

func (s *KillSemanticJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
