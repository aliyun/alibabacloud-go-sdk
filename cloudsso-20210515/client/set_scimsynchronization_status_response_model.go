// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSCIMSynchronizationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSCIMSynchronizationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSCIMSynchronizationStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetSCIMSynchronizationStatusResponseBody) *SetSCIMSynchronizationStatusResponse
	GetBody() *SetSCIMSynchronizationStatusResponseBody
}

type SetSCIMSynchronizationStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSCIMSynchronizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSCIMSynchronizationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSCIMSynchronizationStatusResponse) GoString() string {
	return s.String()
}

func (s *SetSCIMSynchronizationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSCIMSynchronizationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSCIMSynchronizationStatusResponse) GetBody() *SetSCIMSynchronizationStatusResponseBody {
	return s.Body
}

func (s *SetSCIMSynchronizationStatusResponse) SetHeaders(v map[string]*string) *SetSCIMSynchronizationStatusResponse {
	s.Headers = v
	return s
}

func (s *SetSCIMSynchronizationStatusResponse) SetStatusCode(v int32) *SetSCIMSynchronizationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSCIMSynchronizationStatusResponse) SetBody(v *SetSCIMSynchronizationStatusResponseBody) *SetSCIMSynchronizationStatusResponse {
	s.Body = v
	return s
}

func (s *SetSCIMSynchronizationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
