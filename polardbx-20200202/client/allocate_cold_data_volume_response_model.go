// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateColdDataVolumeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateColdDataVolumeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateColdDataVolumeResponse
	GetStatusCode() *int32
	SetBody(v *AllocateColdDataVolumeResponseBody) *AllocateColdDataVolumeResponse
	GetBody() *AllocateColdDataVolumeResponseBody
}

type AllocateColdDataVolumeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateColdDataVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateColdDataVolumeResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateColdDataVolumeResponse) GoString() string {
	return s.String()
}

func (s *AllocateColdDataVolumeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateColdDataVolumeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateColdDataVolumeResponse) GetBody() *AllocateColdDataVolumeResponseBody {
	return s.Body
}

func (s *AllocateColdDataVolumeResponse) SetHeaders(v map[string]*string) *AllocateColdDataVolumeResponse {
	s.Headers = v
	return s
}

func (s *AllocateColdDataVolumeResponse) SetStatusCode(v int32) *AllocateColdDataVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateColdDataVolumeResponse) SetBody(v *AllocateColdDataVolumeResponseBody) *AllocateColdDataVolumeResponse {
	s.Body = v
	return s
}

func (s *AllocateColdDataVolumeResponse) Validate() error {
	return dara.Validate(s)
}
