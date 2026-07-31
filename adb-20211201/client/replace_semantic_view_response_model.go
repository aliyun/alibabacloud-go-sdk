// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceSemanticViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceSemanticViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceSemanticViewResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceSemanticViewResponseBody) *ReplaceSemanticViewResponse
	GetBody() *ReplaceSemanticViewResponseBody
}

type ReplaceSemanticViewResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceSemanticViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceSemanticViewResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceSemanticViewResponse) GoString() string {
	return s.String()
}

func (s *ReplaceSemanticViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceSemanticViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceSemanticViewResponse) GetBody() *ReplaceSemanticViewResponseBody {
	return s.Body
}

func (s *ReplaceSemanticViewResponse) SetHeaders(v map[string]*string) *ReplaceSemanticViewResponse {
	s.Headers = v
	return s
}

func (s *ReplaceSemanticViewResponse) SetStatusCode(v int32) *ReplaceSemanticViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceSemanticViewResponse) SetBody(v *ReplaceSemanticViewResponseBody) *ReplaceSemanticViewResponse {
	s.Body = v
	return s
}

func (s *ReplaceSemanticViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
