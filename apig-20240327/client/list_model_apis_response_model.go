// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListModelApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListModelApisResponse
	GetStatusCode() *int32
	SetBody(v *ListModelApisResponseBody) *ListModelApisResponse
	GetBody() *ListModelApisResponseBody
}

type ListModelApisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelApisResponse) String() string {
	return dara.Prettify(s)
}

func (s ListModelApisResponse) GoString() string {
	return s.String()
}

func (s *ListModelApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListModelApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListModelApisResponse) GetBody() *ListModelApisResponseBody {
	return s.Body
}

func (s *ListModelApisResponse) SetHeaders(v map[string]*string) *ListModelApisResponse {
	s.Headers = v
	return s
}

func (s *ListModelApisResponse) SetStatusCode(v int32) *ListModelApisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelApisResponse) SetBody(v *ListModelApisResponseBody) *ListModelApisResponse {
	s.Body = v
	return s
}

func (s *ListModelApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
