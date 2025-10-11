// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplicationPolarFSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachApplicationPolarFSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachApplicationPolarFSResponse
	GetStatusCode() *int32
	SetBody(v *AttachApplicationPolarFSResponseBody) *AttachApplicationPolarFSResponse
	GetBody() *AttachApplicationPolarFSResponseBody
}

type AttachApplicationPolarFSResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachApplicationPolarFSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachApplicationPolarFSResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachApplicationPolarFSResponse) GoString() string {
	return s.String()
}

func (s *AttachApplicationPolarFSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachApplicationPolarFSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachApplicationPolarFSResponse) GetBody() *AttachApplicationPolarFSResponseBody {
	return s.Body
}

func (s *AttachApplicationPolarFSResponse) SetHeaders(v map[string]*string) *AttachApplicationPolarFSResponse {
	s.Headers = v
	return s
}

func (s *AttachApplicationPolarFSResponse) SetStatusCode(v int32) *AttachApplicationPolarFSResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachApplicationPolarFSResponse) SetBody(v *AttachApplicationPolarFSResponseBody) *AttachApplicationPolarFSResponse {
	s.Body = v
	return s
}

func (s *AttachApplicationPolarFSResponse) Validate() error {
	return dara.Validate(s)
}
