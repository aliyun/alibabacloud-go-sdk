// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkReplSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkReplSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkReplSessionResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkReplSessionResponseBody) *GetSparkReplSessionResponse
	GetBody() *GetSparkReplSessionResponseBody
}

type GetSparkReplSessionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkReplSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkReplSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkReplSessionResponse) GoString() string {
	return s.String()
}

func (s *GetSparkReplSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkReplSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkReplSessionResponse) GetBody() *GetSparkReplSessionResponseBody {
	return s.Body
}

func (s *GetSparkReplSessionResponse) SetHeaders(v map[string]*string) *GetSparkReplSessionResponse {
	s.Headers = v
	return s
}

func (s *GetSparkReplSessionResponse) SetStatusCode(v int32) *GetSparkReplSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkReplSessionResponse) SetBody(v *GetSparkReplSessionResponseBody) *GetSparkReplSessionResponse {
	s.Body = v
	return s
}

func (s *GetSparkReplSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
