// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelOutputContentSyncDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelOutputContentSyncDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelOutputContentSyncDetectResponse
	GetStatusCode() *int32
	SetBody(v *ModelOutputContentSyncDetectResponseBody) *ModelOutputContentSyncDetectResponse
	GetBody() *ModelOutputContentSyncDetectResponseBody
}

type ModelOutputContentSyncDetectResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelOutputContentSyncDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelOutputContentSyncDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelOutputContentSyncDetectResponse) GoString() string {
	return s.String()
}

func (s *ModelOutputContentSyncDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelOutputContentSyncDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelOutputContentSyncDetectResponse) GetBody() *ModelOutputContentSyncDetectResponseBody {
	return s.Body
}

func (s *ModelOutputContentSyncDetectResponse) SetHeaders(v map[string]*string) *ModelOutputContentSyncDetectResponse {
	s.Headers = v
	return s
}

func (s *ModelOutputContentSyncDetectResponse) SetStatusCode(v int32) *ModelOutputContentSyncDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelOutputContentSyncDetectResponse) SetBody(v *ModelOutputContentSyncDetectResponseBody) *ModelOutputContentSyncDetectResponse {
	s.Body = v
	return s
}

func (s *ModelOutputContentSyncDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
