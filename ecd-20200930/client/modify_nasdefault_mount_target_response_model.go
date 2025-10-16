// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNASDefaultMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyNASDefaultMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyNASDefaultMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyNASDefaultMountTargetResponseBody) *ModifyNASDefaultMountTargetResponse
	GetBody() *ModifyNASDefaultMountTargetResponseBody
}

type ModifyNASDefaultMountTargetResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNASDefaultMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNASDefaultMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ModifyNASDefaultMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyNASDefaultMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyNASDefaultMountTargetResponse) GetBody() *ModifyNASDefaultMountTargetResponseBody {
	return s.Body
}

func (s *ModifyNASDefaultMountTargetResponse) SetHeaders(v map[string]*string) *ModifyNASDefaultMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ModifyNASDefaultMountTargetResponse) SetStatusCode(v int32) *ModifyNASDefaultMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNASDefaultMountTargetResponse) SetBody(v *ModifyNASDefaultMountTargetResponseBody) *ModifyNASDefaultMountTargetResponse {
	s.Body = v
	return s
}

func (s *ModifyNASDefaultMountTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
