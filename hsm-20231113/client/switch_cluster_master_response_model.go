// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchClusterMasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchClusterMasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchClusterMasterResponse
	GetStatusCode() *int32
	SetBody(v *SwitchClusterMasterResponseBody) *SwitchClusterMasterResponse
	GetBody() *SwitchClusterMasterResponseBody
}

type SwitchClusterMasterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchClusterMasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchClusterMasterResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchClusterMasterResponse) GoString() string {
	return s.String()
}

func (s *SwitchClusterMasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchClusterMasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchClusterMasterResponse) GetBody() *SwitchClusterMasterResponseBody {
	return s.Body
}

func (s *SwitchClusterMasterResponse) SetHeaders(v map[string]*string) *SwitchClusterMasterResponse {
	s.Headers = v
	return s
}

func (s *SwitchClusterMasterResponse) SetStatusCode(v int32) *SwitchClusterMasterResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchClusterMasterResponse) SetBody(v *SwitchClusterMasterResponseBody) *SwitchClusterMasterResponse {
	s.Body = v
	return s
}

func (s *SwitchClusterMasterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
