// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextsResponse
	GetStatusCode() *int32
	SetBody(v *TextQueryResult) *ListTextsResponse
	GetBody() *TextQueryResult
}

type ListTextsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextQueryResult   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextsResponse) GoString() string {
	return s.String()
}

func (s *ListTextsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextsResponse) GetBody() *TextQueryResult {
	return s.Body
}

func (s *ListTextsResponse) SetHeaders(v map[string]*string) *ListTextsResponse {
	s.Headers = v
	return s
}

func (s *ListTextsResponse) SetStatusCode(v int32) *ListTextsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextsResponse) SetBody(v *TextQueryResult) *ListTextsResponse {
	s.Body = v
	return s
}

func (s *ListTextsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
