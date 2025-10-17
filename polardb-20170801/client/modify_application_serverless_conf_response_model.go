// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationServerlessConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApplicationServerlessConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApplicationServerlessConfResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApplicationServerlessConfResponseBody) *ModifyApplicationServerlessConfResponse
	GetBody() *ModifyApplicationServerlessConfResponseBody
}

type ModifyApplicationServerlessConfResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationServerlessConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationServerlessConfResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationServerlessConfResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationServerlessConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApplicationServerlessConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApplicationServerlessConfResponse) GetBody() *ModifyApplicationServerlessConfResponseBody {
	return s.Body
}

func (s *ModifyApplicationServerlessConfResponse) SetHeaders(v map[string]*string) *ModifyApplicationServerlessConfResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationServerlessConfResponse) SetStatusCode(v int32) *ModifyApplicationServerlessConfResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationServerlessConfResponse) SetBody(v *ModifyApplicationServerlessConfResponseBody) *ModifyApplicationServerlessConfResponse {
	s.Body = v
	return s
}

func (s *ModifyApplicationServerlessConfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
