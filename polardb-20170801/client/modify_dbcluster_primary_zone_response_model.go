// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterPrimaryZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterPrimaryZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterPrimaryZoneResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterPrimaryZoneResponseBody) *ModifyDBClusterPrimaryZoneResponse
	GetBody() *ModifyDBClusterPrimaryZoneResponseBody
}

type ModifyDBClusterPrimaryZoneResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterPrimaryZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterPrimaryZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterPrimaryZoneResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPrimaryZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterPrimaryZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterPrimaryZoneResponse) GetBody() *ModifyDBClusterPrimaryZoneResponseBody {
	return s.Body
}

func (s *ModifyDBClusterPrimaryZoneResponse) SetHeaders(v map[string]*string) *ModifyDBClusterPrimaryZoneResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterPrimaryZoneResponse) SetStatusCode(v int32) *ModifyDBClusterPrimaryZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneResponse) SetBody(v *ModifyDBClusterPrimaryZoneResponseBody) *ModifyDBClusterPrimaryZoneResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterPrimaryZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
