// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAndroidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAndroidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAndroidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAndroidInstanceResponseBody) *ModifyAndroidInstanceResponse
	GetBody() *ModifyAndroidInstanceResponseBody
}

type ModifyAndroidInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAndroidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAndroidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAndroidInstanceResponse) GetBody() *ModifyAndroidInstanceResponseBody {
	return s.Body
}

func (s *ModifyAndroidInstanceResponse) SetHeaders(v map[string]*string) *ModifyAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyAndroidInstanceResponse) SetStatusCode(v int32) *ModifyAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAndroidInstanceResponse) SetBody(v *ModifyAndroidInstanceResponseBody) *ModifyAndroidInstanceResponse {
	s.Body = v
	return s
}

func (s *ModifyAndroidInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
