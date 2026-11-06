// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuilderImageByParentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuilderImageByParentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuilderImageByParentResponse
	GetStatusCode() *int32
	SetBody(v *BuilderImageByParentResponseBody) *BuilderImageByParentResponse
	GetBody() *BuilderImageByParentResponseBody
}

type BuilderImageByParentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuilderImageByParentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuilderImageByParentResponse) String() string {
	return dara.Prettify(s)
}

func (s BuilderImageByParentResponse) GoString() string {
	return s.String()
}

func (s *BuilderImageByParentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuilderImageByParentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuilderImageByParentResponse) GetBody() *BuilderImageByParentResponseBody {
	return s.Body
}

func (s *BuilderImageByParentResponse) SetHeaders(v map[string]*string) *BuilderImageByParentResponse {
	s.Headers = v
	return s
}

func (s *BuilderImageByParentResponse) SetStatusCode(v int32) *BuilderImageByParentResponse {
	s.StatusCode = &v
	return s
}

func (s *BuilderImageByParentResponse) SetBody(v *BuilderImageByParentResponseBody) *BuilderImageByParentResponse {
	s.Body = v
	return s
}

func (s *BuilderImageByParentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
