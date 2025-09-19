// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUninstallClientsByUuidsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUninstallClientsByUuidsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUninstallClientsByUuidsResponse
	GetStatusCode() *int32
	SetBody(v *AddUninstallClientsByUuidsResponseBody) *AddUninstallClientsByUuidsResponse
	GetBody() *AddUninstallClientsByUuidsResponseBody
}

type AddUninstallClientsByUuidsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUninstallClientsByUuidsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUninstallClientsByUuidsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUninstallClientsByUuidsResponse) GoString() string {
	return s.String()
}

func (s *AddUninstallClientsByUuidsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUninstallClientsByUuidsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUninstallClientsByUuidsResponse) GetBody() *AddUninstallClientsByUuidsResponseBody {
	return s.Body
}

func (s *AddUninstallClientsByUuidsResponse) SetHeaders(v map[string]*string) *AddUninstallClientsByUuidsResponse {
	s.Headers = v
	return s
}

func (s *AddUninstallClientsByUuidsResponse) SetStatusCode(v int32) *AddUninstallClientsByUuidsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUninstallClientsByUuidsResponse) SetBody(v *AddUninstallClientsByUuidsResponseBody) *AddUninstallClientsByUuidsResponse {
	s.Body = v
	return s
}

func (s *AddUninstallClientsByUuidsResponse) Validate() error {
	return dara.Validate(s)
}
