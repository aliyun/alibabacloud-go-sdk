// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDrdsInstanceDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDrdsInstanceDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDrdsInstanceDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDrdsInstanceDescriptionResponseBody) *ModifyDrdsInstanceDescriptionResponse
	GetBody() *ModifyDrdsInstanceDescriptionResponseBody
}

type ModifyDrdsInstanceDescriptionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDrdsInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDrdsInstanceDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDrdsInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDrdsInstanceDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDrdsInstanceDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDrdsInstanceDescriptionResponse) GetBody() *ModifyDrdsInstanceDescriptionResponseBody {
	return s.Body
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDrdsInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDrdsInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponse) SetBody(v *ModifyDrdsInstanceDescriptionResponseBody) *ModifyDrdsInstanceDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifyDrdsInstanceDescriptionResponse) Validate() error {
	return dara.Validate(s)
}
