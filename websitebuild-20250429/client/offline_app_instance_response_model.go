// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAppInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineAppInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineAppInstanceResponse
	GetStatusCode() *int32
	SetBody(v *OfflineAppInstanceResponseBody) *OfflineAppInstanceResponse
	GetBody() *OfflineAppInstanceResponseBody
}

type OfflineAppInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineAppInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineAppInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineAppInstanceResponse) GoString() string {
	return s.String()
}

func (s *OfflineAppInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineAppInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineAppInstanceResponse) GetBody() *OfflineAppInstanceResponseBody {
	return s.Body
}

func (s *OfflineAppInstanceResponse) SetHeaders(v map[string]*string) *OfflineAppInstanceResponse {
	s.Headers = v
	return s
}

func (s *OfflineAppInstanceResponse) SetStatusCode(v int32) *OfflineAppInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineAppInstanceResponse) SetBody(v *OfflineAppInstanceResponseBody) *OfflineAppInstanceResponse {
	s.Body = v
	return s
}

func (s *OfflineAppInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
