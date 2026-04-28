// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStoryResponse
	GetStatusCode() *int32
	SetBody(v *Story) *GetStoryResponse
	GetBody() *Story
}

type GetStoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Story             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStoryResponse) GoString() string {
	return s.String()
}

func (s *GetStoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStoryResponse) GetBody() *Story {
	return s.Body
}

func (s *GetStoryResponse) SetHeaders(v map[string]*string) *GetStoryResponse {
	s.Headers = v
	return s
}

func (s *GetStoryResponse) SetStatusCode(v int32) *GetStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStoryResponse) SetBody(v *Story) *GetStoryResponse {
	s.Body = v
	return s
}

func (s *GetStoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
