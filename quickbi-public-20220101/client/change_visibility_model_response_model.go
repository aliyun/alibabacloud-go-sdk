// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVisibilityModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeVisibilityModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeVisibilityModelResponse
	GetStatusCode() *int32
	SetBody(v *ChangeVisibilityModelResponseBody) *ChangeVisibilityModelResponse
	GetBody() *ChangeVisibilityModelResponseBody
}

type ChangeVisibilityModelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeVisibilityModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeVisibilityModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeVisibilityModelResponse) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeVisibilityModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeVisibilityModelResponse) GetBody() *ChangeVisibilityModelResponseBody {
	return s.Body
}

func (s *ChangeVisibilityModelResponse) SetHeaders(v map[string]*string) *ChangeVisibilityModelResponse {
	s.Headers = v
	return s
}

func (s *ChangeVisibilityModelResponse) SetStatusCode(v int32) *ChangeVisibilityModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeVisibilityModelResponse) SetBody(v *ChangeVisibilityModelResponseBody) *ChangeVisibilityModelResponse {
	s.Body = v
	return s
}

func (s *ChangeVisibilityModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
