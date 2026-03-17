// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeSmartAGWebConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SynchronizeSmartAGWebConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SynchronizeSmartAGWebConfigResponse
	GetStatusCode() *int32
	SetBody(v *SynchronizeSmartAGWebConfigResponseBody) *SynchronizeSmartAGWebConfigResponse
	GetBody() *SynchronizeSmartAGWebConfigResponseBody
}

type SynchronizeSmartAGWebConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynchronizeSmartAGWebConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SynchronizeSmartAGWebConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeSmartAGWebConfigResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeSmartAGWebConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SynchronizeSmartAGWebConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SynchronizeSmartAGWebConfigResponse) GetBody() *SynchronizeSmartAGWebConfigResponseBody {
	return s.Body
}

func (s *SynchronizeSmartAGWebConfigResponse) SetHeaders(v map[string]*string) *SynchronizeSmartAGWebConfigResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeSmartAGWebConfigResponse) SetStatusCode(v int32) *SynchronizeSmartAGWebConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigResponse) SetBody(v *SynchronizeSmartAGWebConfigResponseBody) *SynchronizeSmartAGWebConfigResponse {
	s.Body = v
	return s
}

func (s *SynchronizeSmartAGWebConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
