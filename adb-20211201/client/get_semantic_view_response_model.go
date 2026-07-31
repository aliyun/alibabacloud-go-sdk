// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSemanticViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSemanticViewResponse
	GetStatusCode() *int32
	SetBody(v *GetSemanticViewResponseBody) *GetSemanticViewResponse
	GetBody() *GetSemanticViewResponseBody
}

type GetSemanticViewResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSemanticViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSemanticViewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticViewResponse) GoString() string {
	return s.String()
}

func (s *GetSemanticViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSemanticViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSemanticViewResponse) GetBody() *GetSemanticViewResponseBody {
	return s.Body
}

func (s *GetSemanticViewResponse) SetHeaders(v map[string]*string) *GetSemanticViewResponse {
	s.Headers = v
	return s
}

func (s *GetSemanticViewResponse) SetStatusCode(v int32) *GetSemanticViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSemanticViewResponse) SetBody(v *GetSemanticViewResponseBody) *GetSemanticViewResponse {
	s.Body = v
	return s
}

func (s *GetSemanticViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
