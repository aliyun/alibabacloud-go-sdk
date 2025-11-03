// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVbrToVpconnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachVbrToVpconnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachVbrToVpconnResponse
	GetStatusCode() *int32
	SetBody(v *AttachVbrToVpconnResponseBody) *AttachVbrToVpconnResponse
	GetBody() *AttachVbrToVpconnResponseBody
}

type AttachVbrToVpconnResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachVbrToVpconnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachVbrToVpconnResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachVbrToVpconnResponse) GoString() string {
	return s.String()
}

func (s *AttachVbrToVpconnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachVbrToVpconnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachVbrToVpconnResponse) GetBody() *AttachVbrToVpconnResponseBody {
	return s.Body
}

func (s *AttachVbrToVpconnResponse) SetHeaders(v map[string]*string) *AttachVbrToVpconnResponse {
	s.Headers = v
	return s
}

func (s *AttachVbrToVpconnResponse) SetStatusCode(v int32) *AttachVbrToVpconnResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachVbrToVpconnResponse) SetBody(v *AttachVbrToVpconnResponseBody) *AttachVbrToVpconnResponse {
	s.Body = v
	return s
}

func (s *AttachVbrToVpconnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
