// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDefenseResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDefenseResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDefenseResourceResponseBody) *CreateDefenseResourceResponse
	GetBody() *CreateDefenseResourceResponseBody
}

type CreateDefenseResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefenseResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefenseResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDefenseResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDefenseResourceResponse) GetBody() *CreateDefenseResourceResponseBody {
	return s.Body
}

func (s *CreateDefenseResourceResponse) SetHeaders(v map[string]*string) *CreateDefenseResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDefenseResourceResponse) SetStatusCode(v int32) *CreateDefenseResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefenseResourceResponse) SetBody(v *CreateDefenseResourceResponseBody) *CreateDefenseResourceResponse {
	s.Body = v
	return s
}

func (s *CreateDefenseResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
