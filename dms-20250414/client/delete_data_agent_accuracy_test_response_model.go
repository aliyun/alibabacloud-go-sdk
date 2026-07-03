// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentAccuracyTestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataAgentAccuracyTestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataAgentAccuracyTestResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataAgentAccuracyTestResponseBody) *DeleteDataAgentAccuracyTestResponse
	GetBody() *DeleteDataAgentAccuracyTestResponseBody
}

type DeleteDataAgentAccuracyTestResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataAgentAccuracyTestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataAgentAccuracyTestResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentAccuracyTestResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentAccuracyTestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataAgentAccuracyTestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataAgentAccuracyTestResponse) GetBody() *DeleteDataAgentAccuracyTestResponseBody {
	return s.Body
}

func (s *DeleteDataAgentAccuracyTestResponse) SetHeaders(v map[string]*string) *DeleteDataAgentAccuracyTestResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataAgentAccuracyTestResponse) SetStatusCode(v int32) *DeleteDataAgentAccuracyTestResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataAgentAccuracyTestResponse) SetBody(v *DeleteDataAgentAccuracyTestResponseBody) *DeleteDataAgentAccuracyTestResponse {
	s.Body = v
	return s
}

func (s *DeleteDataAgentAccuracyTestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
