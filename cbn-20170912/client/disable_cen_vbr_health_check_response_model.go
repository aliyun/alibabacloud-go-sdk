// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCenVbrHealthCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCenVbrHealthCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCenVbrHealthCheckResponse
	GetStatusCode() *int32
	SetBody(v *DisableCenVbrHealthCheckResponseBody) *DisableCenVbrHealthCheckResponse
	GetBody() *DisableCenVbrHealthCheckResponseBody
}

type DisableCenVbrHealthCheckResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCenVbrHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCenVbrHealthCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCenVbrHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *DisableCenVbrHealthCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCenVbrHealthCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCenVbrHealthCheckResponse) GetBody() *DisableCenVbrHealthCheckResponseBody {
	return s.Body
}

func (s *DisableCenVbrHealthCheckResponse) SetHeaders(v map[string]*string) *DisableCenVbrHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *DisableCenVbrHealthCheckResponse) SetStatusCode(v int32) *DisableCenVbrHealthCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCenVbrHealthCheckResponse) SetBody(v *DisableCenVbrHealthCheckResponseBody) *DisableCenVbrHealthCheckResponse {
	s.Body = v
	return s
}

func (s *DisableCenVbrHealthCheckResponse) Validate() error {
	return dara.Validate(s)
}
