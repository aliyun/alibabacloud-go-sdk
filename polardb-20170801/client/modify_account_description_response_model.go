// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccountDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse
	GetBody() *ModifyAccountDescriptionResponseBody
}

type ModifyAccountDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccountDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccountDescriptionResponse) GetBody() *ModifyAccountDescriptionResponseBody {
	return s.Body
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetStatusCode(v int32) *ModifyAccountDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyAccountDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
