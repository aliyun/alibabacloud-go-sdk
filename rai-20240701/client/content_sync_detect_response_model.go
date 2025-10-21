// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContentSyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContentSyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContentSyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *ContentSyncDetectResponseBody) *ContentSyncDetectResponse
	GetBody() *ContentSyncDetectResponseBody
}

type ContentSyncDetectResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContentSyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s ContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ContentSyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContentSyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContentSyncDetectResponse) GetBody() *ContentSyncDetectResponseBody {
	return s.Body
}

func (s *ContentSyncDetectResponse) SetHeaders(v map[string]*string) *ContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ContentSyncDetectResponse) SetStatusCode(v int32) *ContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ContentSyncDetectResponse) SetBody(v *ContentSyncDetectResponseBody) *ContentSyncDetectResponse {
	s.Body = v
	return s
}

func (s *ContentSyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
