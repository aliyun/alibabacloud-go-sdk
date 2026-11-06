// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarFsInstanceDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolarFsInstanceDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolarFsInstanceDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolarFsInstanceDescriptionResponseBody) *ModifyPolarFsInstanceDescriptionResponse
	GetBody() *ModifyPolarFsInstanceDescriptionResponseBody
}

type ModifyPolarFsInstanceDescriptionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolarFsInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolarFsInstanceDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarFsInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolarFsInstanceDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolarFsInstanceDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolarFsInstanceDescriptionResponse) GetBody() *ModifyPolarFsInstanceDescriptionResponseBody {
	return s.Body
}

func (s *ModifyPolarFsInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyPolarFsInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyPolarFsInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionResponse) SetBody(v *ModifyPolarFsInstanceDescriptionResponseBody) *ModifyPolarFsInstanceDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyPolarFsInstanceDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
