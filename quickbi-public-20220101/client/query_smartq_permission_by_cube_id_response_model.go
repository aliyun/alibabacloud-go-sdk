// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmartqPermissionByCubeIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySmartqPermissionByCubeIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySmartqPermissionByCubeIdResponse
	GetStatusCode() *int32
	SetBody(v *QuerySmartqPermissionByCubeIdResponseBody) *QuerySmartqPermissionByCubeIdResponse
	GetBody() *QuerySmartqPermissionByCubeIdResponseBody
}

type QuerySmartqPermissionByCubeIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySmartqPermissionByCubeIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySmartqPermissionByCubeIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySmartqPermissionByCubeIdResponse) GoString() string {
	return s.String()
}

func (s *QuerySmartqPermissionByCubeIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySmartqPermissionByCubeIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySmartqPermissionByCubeIdResponse) GetBody() *QuerySmartqPermissionByCubeIdResponseBody {
	return s.Body
}

func (s *QuerySmartqPermissionByCubeIdResponse) SetHeaders(v map[string]*string) *QuerySmartqPermissionByCubeIdResponse {
	s.Headers = v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponse) SetStatusCode(v int32) *QuerySmartqPermissionByCubeIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponse) SetBody(v *QuerySmartqPermissionByCubeIdResponseBody) *QuerySmartqPermissionByCubeIdResponse {
	s.Body = v
	return s
}

func (s *QuerySmartqPermissionByCubeIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
