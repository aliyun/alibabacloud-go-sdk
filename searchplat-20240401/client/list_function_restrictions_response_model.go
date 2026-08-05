// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionRestrictionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFunctionRestrictionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFunctionRestrictionsResponse
	GetStatusCode() *int32
	SetBody(v *ListFunctionRestrictionsResponseBody) *ListFunctionRestrictionsResponse
	GetBody() *ListFunctionRestrictionsResponseBody
}

type ListFunctionRestrictionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionRestrictionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionRestrictionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionRestrictionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionRestrictionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFunctionRestrictionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFunctionRestrictionsResponse) GetBody() *ListFunctionRestrictionsResponseBody {
	return s.Body
}

func (s *ListFunctionRestrictionsResponse) SetHeaders(v map[string]*string) *ListFunctionRestrictionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionRestrictionsResponse) SetStatusCode(v int32) *ListFunctionRestrictionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionRestrictionsResponse) SetBody(v *ListFunctionRestrictionsResponseBody) *ListFunctionRestrictionsResponse {
	s.Body = v
	return s
}

func (s *ListFunctionRestrictionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
