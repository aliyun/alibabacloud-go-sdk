// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncFlowResponse
	GetStatusCode() *int32
	SetBody(v *SyncFlowResponseBody) *SyncFlowResponse
	GetBody() *SyncFlowResponseBody
}

type SyncFlowResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncFlowResponse) GoString() string {
	return s.String()
}

func (s *SyncFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncFlowResponse) GetBody() *SyncFlowResponseBody {
	return s.Body
}

func (s *SyncFlowResponse) SetHeaders(v map[string]*string) *SyncFlowResponse {
	s.Headers = v
	return s
}

func (s *SyncFlowResponse) SetStatusCode(v int32) *SyncFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncFlowResponse) SetBody(v *SyncFlowResponseBody) *SyncFlowResponse {
	s.Body = v
	return s
}

func (s *SyncFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
