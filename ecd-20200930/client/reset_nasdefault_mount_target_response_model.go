// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetNASDefaultMountTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetNASDefaultMountTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetNASDefaultMountTargetResponse
	GetStatusCode() *int32
	SetBody(v *ResetNASDefaultMountTargetResponseBody) *ResetNASDefaultMountTargetResponse
	GetBody() *ResetNASDefaultMountTargetResponseBody
}

type ResetNASDefaultMountTargetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetNASDefaultMountTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetNASDefaultMountTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetNASDefaultMountTargetResponse) GoString() string {
	return s.String()
}

func (s *ResetNASDefaultMountTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetNASDefaultMountTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetNASDefaultMountTargetResponse) GetBody() *ResetNASDefaultMountTargetResponseBody {
	return s.Body
}

func (s *ResetNASDefaultMountTargetResponse) SetHeaders(v map[string]*string) *ResetNASDefaultMountTargetResponse {
	s.Headers = v
	return s
}

func (s *ResetNASDefaultMountTargetResponse) SetStatusCode(v int32) *ResetNASDefaultMountTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetNASDefaultMountTargetResponse) SetBody(v *ResetNASDefaultMountTargetResponseBody) *ResetNASDefaultMountTargetResponse {
	s.Body = v
	return s
}

func (s *ResetNASDefaultMountTargetResponse) Validate() error {
	return dara.Validate(s)
}
