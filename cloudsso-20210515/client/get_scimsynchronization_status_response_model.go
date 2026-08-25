// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSCIMSynchronizationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSCIMSynchronizationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSCIMSynchronizationStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetSCIMSynchronizationStatusResponseBody) *GetSCIMSynchronizationStatusResponse
	GetBody() *GetSCIMSynchronizationStatusResponseBody
}

type GetSCIMSynchronizationStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSCIMSynchronizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSCIMSynchronizationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSCIMSynchronizationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetSCIMSynchronizationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSCIMSynchronizationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSCIMSynchronizationStatusResponse) GetBody() *GetSCIMSynchronizationStatusResponseBody {
	return s.Body
}

func (s *GetSCIMSynchronizationStatusResponse) SetHeaders(v map[string]*string) *GetSCIMSynchronizationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetSCIMSynchronizationStatusResponse) SetStatusCode(v int32) *GetSCIMSynchronizationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSCIMSynchronizationStatusResponse) SetBody(v *GetSCIMSynchronizationStatusResponseBody) *GetSCIMSynchronizationStatusResponse {
	s.Body = v
	return s
}

func (s *GetSCIMSynchronizationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
