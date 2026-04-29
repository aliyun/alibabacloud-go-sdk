// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchPolarClawConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PatchPolarClawConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PatchPolarClawConfigResponse
	GetStatusCode() *int32
	SetBody(v *PatchPolarClawConfigResponseBody) *PatchPolarClawConfigResponse
	GetBody() *PatchPolarClawConfigResponseBody
}

type PatchPolarClawConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PatchPolarClawConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PatchPolarClawConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s PatchPolarClawConfigResponse) GoString() string {
	return s.String()
}

func (s *PatchPolarClawConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PatchPolarClawConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PatchPolarClawConfigResponse) GetBody() *PatchPolarClawConfigResponseBody {
	return s.Body
}

func (s *PatchPolarClawConfigResponse) SetHeaders(v map[string]*string) *PatchPolarClawConfigResponse {
	s.Headers = v
	return s
}

func (s *PatchPolarClawConfigResponse) SetStatusCode(v int32) *PatchPolarClawConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *PatchPolarClawConfigResponse) SetBody(v *PatchPolarClawConfigResponseBody) *PatchPolarClawConfigResponse {
	s.Body = v
	return s
}

func (s *PatchPolarClawConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
