// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicInfoQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBasicInfoQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBasicInfoQAResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBasicInfoQAResponseBody) *UpdateBasicInfoQAResponse
	GetBody() *UpdateBasicInfoQAResponseBody
}

type UpdateBasicInfoQAResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBasicInfoQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBasicInfoQAResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicInfoQAResponse) GoString() string {
	return s.String()
}

func (s *UpdateBasicInfoQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBasicInfoQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBasicInfoQAResponse) GetBody() *UpdateBasicInfoQAResponseBody {
	return s.Body
}

func (s *UpdateBasicInfoQAResponse) SetHeaders(v map[string]*string) *UpdateBasicInfoQAResponse {
	s.Headers = v
	return s
}

func (s *UpdateBasicInfoQAResponse) SetStatusCode(v int32) *UpdateBasicInfoQAResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBasicInfoQAResponse) SetBody(v *UpdateBasicInfoQAResponseBody) *UpdateBasicInfoQAResponse {
	s.Body = v
	return s
}

func (s *UpdateBasicInfoQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
