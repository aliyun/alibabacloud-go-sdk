// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationDescriptionResponseBody) *ModifyApplicationDescriptionResponse
	GetBody() *ModifyApplicationDescriptionResponseBody
}

type ModifyApplicationDescriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationDescriptionResponse) GetBody() *ModifyApplicationDescriptionResponseBody {
	return s.Body
}

func (s *ModifyApplicationDescriptionResponse) SetHeaders(v map[string]*string) *ModifyApplicationDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationDescriptionResponse) SetStatusCode(v int32) *ModifyApplicationDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationDescriptionResponse) SetBody(v *ModifyApplicationDescriptionResponseBody) *ModifyApplicationDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
