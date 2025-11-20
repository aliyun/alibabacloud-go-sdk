// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrieveRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrieveRunResponse
	GetStatusCode() *int32
	SetBody(v *RetrieveRunResponseBody) *RetrieveRunResponse
	GetBody() *RetrieveRunResponseBody
}

type RetrieveRunResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveRunResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRunResponse) GoString() string {
	return s.String()
}

func (s *RetrieveRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrieveRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrieveRunResponse) GetBody() *RetrieveRunResponseBody {
	return s.Body
}

func (s *RetrieveRunResponse) SetHeaders(v map[string]*string) *RetrieveRunResponse {
	s.Headers = v
	return s
}

func (s *RetrieveRunResponse) SetStatusCode(v int32) *RetrieveRunResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveRunResponse) SetBody(v *RetrieveRunResponseBody) *RetrieveRunResponse {
	s.Body = v
	return s
}

func (s *RetrieveRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
