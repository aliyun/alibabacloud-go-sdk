// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseLogSyncToSLSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseLogSyncToSLSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseLogSyncToSLSResponse
	GetStatusCode() *int32
	SetBody(v *CloseLogSyncToSLSResponseBody) *CloseLogSyncToSLSResponse
	GetBody() *CloseLogSyncToSLSResponseBody
}

type CloseLogSyncToSLSResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseLogSyncToSLSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseLogSyncToSLSResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseLogSyncToSLSResponse) GoString() string {
	return s.String()
}

func (s *CloseLogSyncToSLSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseLogSyncToSLSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseLogSyncToSLSResponse) GetBody() *CloseLogSyncToSLSResponseBody {
	return s.Body
}

func (s *CloseLogSyncToSLSResponse) SetHeaders(v map[string]*string) *CloseLogSyncToSLSResponse {
	s.Headers = v
	return s
}

func (s *CloseLogSyncToSLSResponse) SetStatusCode(v int32) *CloseLogSyncToSLSResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseLogSyncToSLSResponse) SetBody(v *CloseLogSyncToSLSResponseBody) *CloseLogSyncToSLSResponse {
	s.Body = v
	return s
}

func (s *CloseLogSyncToSLSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
