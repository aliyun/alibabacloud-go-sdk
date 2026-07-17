// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeEditingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateYikeEditingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateYikeEditingProjectResponse
	GetStatusCode() *int32
	SetBody(v *CreateYikeEditingProjectResponseBody) *CreateYikeEditingProjectResponse
	GetBody() *CreateYikeEditingProjectResponseBody
}

type CreateYikeEditingProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateYikeEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateYikeEditingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateYikeEditingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateYikeEditingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateYikeEditingProjectResponse) GetBody() *CreateYikeEditingProjectResponseBody {
	return s.Body
}

func (s *CreateYikeEditingProjectResponse) SetHeaders(v map[string]*string) *CreateYikeEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateYikeEditingProjectResponse) SetStatusCode(v int32) *CreateYikeEditingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateYikeEditingProjectResponse) SetBody(v *CreateYikeEditingProjectResponseBody) *CreateYikeEditingProjectResponse {
	s.Body = v
	return s
}

func (s *CreateYikeEditingProjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
