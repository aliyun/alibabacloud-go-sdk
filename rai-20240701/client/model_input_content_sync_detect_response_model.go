// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelInputContentSyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelInputContentSyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelInputContentSyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *ModelInputContentSyncDetectResponseBody) *ModelInputContentSyncDetectResponse
	GetBody() *ModelInputContentSyncDetectResponseBody
}

type ModelInputContentSyncDetectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelInputContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelInputContentSyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelInputContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelInputContentSyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelInputContentSyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelInputContentSyncDetectResponse) GetBody() *ModelInputContentSyncDetectResponseBody {
	return s.Body
}

func (s *ModelInputContentSyncDetectResponse) SetHeaders(v map[string]*string) *ModelInputContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelInputContentSyncDetectResponse) SetStatusCode(v int32) *ModelInputContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelInputContentSyncDetectResponse) SetBody(v *ModelInputContentSyncDetectResponseBody) *ModelInputContentSyncDetectResponse {
	s.Body = v
	return s
}

func (s *ModelInputContentSyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
