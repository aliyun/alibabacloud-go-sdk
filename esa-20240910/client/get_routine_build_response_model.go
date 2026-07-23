// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineBuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineBuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineBuildResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineBuildResponseBody) *GetRoutineBuildResponse
	GetBody() *GetRoutineBuildResponseBody
}

type GetRoutineBuildResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineBuildResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineBuildResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineBuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineBuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineBuildResponse) GetBody() *GetRoutineBuildResponseBody {
	return s.Body
}

func (s *GetRoutineBuildResponse) SetHeaders(v map[string]*string) *GetRoutineBuildResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineBuildResponse) SetStatusCode(v int32) *GetRoutineBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineBuildResponse) SetBody(v *GetRoutineBuildResponseBody) *GetRoutineBuildResponse {
	s.Body = v
	return s
}

func (s *GetRoutineBuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
