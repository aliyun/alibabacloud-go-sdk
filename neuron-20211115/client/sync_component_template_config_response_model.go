// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncComponentTemplateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncComponentTemplateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncComponentTemplateConfigResponse
	GetStatusCode() *int32
	SetBody(v *SyncComponentTemplateConfigResponseBody) *SyncComponentTemplateConfigResponse
	GetBody() *SyncComponentTemplateConfigResponseBody
}

type SyncComponentTemplateConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncComponentTemplateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncComponentTemplateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncComponentTemplateConfigResponse) GoString() string {
	return s.String()
}

func (s *SyncComponentTemplateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncComponentTemplateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncComponentTemplateConfigResponse) GetBody() *SyncComponentTemplateConfigResponseBody {
	return s.Body
}

func (s *SyncComponentTemplateConfigResponse) SetHeaders(v map[string]*string) *SyncComponentTemplateConfigResponse {
	s.Headers = v
	return s
}

func (s *SyncComponentTemplateConfigResponse) SetStatusCode(v int32) *SyncComponentTemplateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncComponentTemplateConfigResponse) SetBody(v *SyncComponentTemplateConfigResponseBody) *SyncComponentTemplateConfigResponse {
	s.Body = v
	return s
}

func (s *SyncComponentTemplateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
