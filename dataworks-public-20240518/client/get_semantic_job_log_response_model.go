// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticJobLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSemanticJobLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSemanticJobLogResponse
	GetStatusCode() *int32
	SetBody(v *GetSemanticJobLogResponseBody) *GetSemanticJobLogResponse
	GetBody() *GetSemanticJobLogResponseBody
}

type GetSemanticJobLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSemanticJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSemanticJobLogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetSemanticJobLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSemanticJobLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSemanticJobLogResponse) GetBody() *GetSemanticJobLogResponseBody {
	return s.Body
}

func (s *GetSemanticJobLogResponse) SetHeaders(v map[string]*string) *GetSemanticJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetSemanticJobLogResponse) SetStatusCode(v int32) *GetSemanticJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSemanticJobLogResponse) SetBody(v *GetSemanticJobLogResponseBody) *GetSemanticJobLogResponse {
	s.Body = v
	return s
}

func (s *GetSemanticJobLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
