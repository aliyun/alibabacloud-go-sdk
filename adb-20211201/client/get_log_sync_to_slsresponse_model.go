// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogSyncToSLSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogSyncToSLSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogSyncToSLSResponse
	GetStatusCode() *int32
	SetBody(v *GetLogSyncToSLSResponseBody) *GetLogSyncToSLSResponse
	GetBody() *GetLogSyncToSLSResponseBody
}

type GetLogSyncToSLSResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogSyncToSLSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogSyncToSLSResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogSyncToSLSResponse) GoString() string {
	return s.String()
}

func (s *GetLogSyncToSLSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogSyncToSLSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogSyncToSLSResponse) GetBody() *GetLogSyncToSLSResponseBody {
	return s.Body
}

func (s *GetLogSyncToSLSResponse) SetHeaders(v map[string]*string) *GetLogSyncToSLSResponse {
	s.Headers = v
	return s
}

func (s *GetLogSyncToSLSResponse) SetStatusCode(v int32) *GetLogSyncToSLSResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogSyncToSLSResponse) SetBody(v *GetLogSyncToSLSResponseBody) *GetLogSyncToSLSResponse {
	s.Body = v
	return s
}

func (s *GetLogSyncToSLSResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
