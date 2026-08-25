// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLogSyncToSLSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenLogSyncToSLSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenLogSyncToSLSResponse
	GetStatusCode() *int32
	SetBody(v *OpenLogSyncToSLSResponseBody) *OpenLogSyncToSLSResponse
	GetBody() *OpenLogSyncToSLSResponseBody
}

type OpenLogSyncToSLSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenLogSyncToSLSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenLogSyncToSLSResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenLogSyncToSLSResponse) GoString() string {
	return s.String()
}

func (s *OpenLogSyncToSLSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenLogSyncToSLSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenLogSyncToSLSResponse) GetBody() *OpenLogSyncToSLSResponseBody {
	return s.Body
}

func (s *OpenLogSyncToSLSResponse) SetHeaders(v map[string]*string) *OpenLogSyncToSLSResponse {
	s.Headers = v
	return s
}

func (s *OpenLogSyncToSLSResponse) SetStatusCode(v int32) *OpenLogSyncToSLSResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenLogSyncToSLSResponse) SetBody(v *OpenLogSyncToSLSResponseBody) *OpenLogSyncToSLSResponse {
	s.Body = v
	return s
}

func (s *OpenLogSyncToSLSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
