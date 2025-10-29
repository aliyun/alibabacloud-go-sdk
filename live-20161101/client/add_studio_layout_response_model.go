// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStudioLayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddStudioLayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddStudioLayoutResponse
	GetStatusCode() *int32
	SetBody(v *AddStudioLayoutResponseBody) *AddStudioLayoutResponse
	GetBody() *AddStudioLayoutResponseBody
}

type AddStudioLayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddStudioLayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddStudioLayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s AddStudioLayoutResponse) GoString() string {
	return s.String()
}

func (s *AddStudioLayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddStudioLayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddStudioLayoutResponse) GetBody() *AddStudioLayoutResponseBody {
	return s.Body
}

func (s *AddStudioLayoutResponse) SetHeaders(v map[string]*string) *AddStudioLayoutResponse {
	s.Headers = v
	return s
}

func (s *AddStudioLayoutResponse) SetStatusCode(v int32) *AddStudioLayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *AddStudioLayoutResponse) SetBody(v *AddStudioLayoutResponseBody) *AddStudioLayoutResponse {
	s.Body = v
	return s
}

func (s *AddStudioLayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
