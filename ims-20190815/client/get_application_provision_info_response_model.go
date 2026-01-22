// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisionInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationProvisionInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationProvisionInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationProvisionInfoResponseBody) *GetApplicationProvisionInfoResponse
	GetBody() *GetApplicationProvisionInfoResponseBody
}

type GetApplicationProvisionInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisionInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationProvisionInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationProvisionInfoResponse) GetBody() *GetApplicationProvisionInfoResponseBody {
	return s.Body
}

func (s *GetApplicationProvisionInfoResponse) SetHeaders(v map[string]*string) *GetApplicationProvisionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisionInfoResponse) SetStatusCode(v int32) *GetApplicationProvisionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisionInfoResponse) SetBody(v *GetApplicationProvisionInfoResponseBody) *GetApplicationProvisionInfoResponse {
	s.Body = v
	return s
}

func (s *GetApplicationProvisionInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
