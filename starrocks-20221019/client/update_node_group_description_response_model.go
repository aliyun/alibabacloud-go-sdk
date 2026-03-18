// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodeGroupDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodeGroupDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodeGroupDescriptionResponseBody) *UpdateNodeGroupDescriptionResponse
	GetBody() *UpdateNodeGroupDescriptionResponseBody
}

type UpdateNodeGroupDescriptionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeGroupDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeGroupDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodeGroupDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodeGroupDescriptionResponse) GetBody() *UpdateNodeGroupDescriptionResponseBody {
	return s.Body
}

func (s *UpdateNodeGroupDescriptionResponse) SetHeaders(v map[string]*string) *UpdateNodeGroupDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeGroupDescriptionResponse) SetStatusCode(v int32) *UpdateNodeGroupDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeGroupDescriptionResponse) SetBody(v *UpdateNodeGroupDescriptionResponseBody) *UpdateNodeGroupDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateNodeGroupDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
