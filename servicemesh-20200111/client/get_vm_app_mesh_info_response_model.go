// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVmAppMeshInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVmAppMeshInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVmAppMeshInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetVmAppMeshInfoResponseBody) *GetVmAppMeshInfoResponse
	GetBody() *GetVmAppMeshInfoResponseBody
}

type GetVmAppMeshInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVmAppMeshInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVmAppMeshInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVmAppMeshInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVmAppMeshInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVmAppMeshInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVmAppMeshInfoResponse) GetBody() *GetVmAppMeshInfoResponseBody {
	return s.Body
}

func (s *GetVmAppMeshInfoResponse) SetHeaders(v map[string]*string) *GetVmAppMeshInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVmAppMeshInfoResponse) SetStatusCode(v int32) *GetVmAppMeshInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVmAppMeshInfoResponse) SetBody(v *GetVmAppMeshInfoResponseBody) *GetVmAppMeshInfoResponse {
	s.Body = v
	return s
}

func (s *GetVmAppMeshInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
