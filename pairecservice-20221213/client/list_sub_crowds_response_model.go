// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubCrowdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubCrowdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubCrowdsResponse
	GetStatusCode() *int32
	SetBody(v *ListSubCrowdsResponseBody) *ListSubCrowdsResponse
	GetBody() *ListSubCrowdsResponseBody
}

type ListSubCrowdsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubCrowdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubCrowdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubCrowdsResponse) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubCrowdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubCrowdsResponse) GetBody() *ListSubCrowdsResponseBody {
	return s.Body
}

func (s *ListSubCrowdsResponse) SetHeaders(v map[string]*string) *ListSubCrowdsResponse {
	s.Headers = v
	return s
}

func (s *ListSubCrowdsResponse) SetStatusCode(v int32) *ListSubCrowdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubCrowdsResponse) SetBody(v *ListSubCrowdsResponseBody) *ListSubCrowdsResponse {
	s.Body = v
	return s
}

func (s *ListSubCrowdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
