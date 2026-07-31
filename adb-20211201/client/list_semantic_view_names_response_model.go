// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticViewNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSemanticViewNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSemanticViewNamesResponse
	GetStatusCode() *int32
	SetBody(v *ListSemanticViewNamesResponseBody) *ListSemanticViewNamesResponse
	GetBody() *ListSemanticViewNamesResponseBody
}

type ListSemanticViewNamesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSemanticViewNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSemanticViewNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticViewNamesResponse) GoString() string {
	return s.String()
}

func (s *ListSemanticViewNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSemanticViewNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSemanticViewNamesResponse) GetBody() *ListSemanticViewNamesResponseBody {
	return s.Body
}

func (s *ListSemanticViewNamesResponse) SetHeaders(v map[string]*string) *ListSemanticViewNamesResponse {
	s.Headers = v
	return s
}

func (s *ListSemanticViewNamesResponse) SetStatusCode(v int32) *ListSemanticViewNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSemanticViewNamesResponse) SetBody(v *ListSemanticViewNamesResponseBody) *ListSemanticViewNamesResponse {
	s.Body = v
	return s
}

func (s *ListSemanticViewNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
