// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTopicSelectionMergeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTopicSelectionMergeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTopicSelectionMergeResponse
	GetStatusCode() *int32
	SetBody(v *RunTopicSelectionMergeResponseBody) *RunTopicSelectionMergeResponse
	GetBody() *RunTopicSelectionMergeResponseBody
}

type RunTopicSelectionMergeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTopicSelectionMergeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTopicSelectionMergeResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeResponse) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTopicSelectionMergeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTopicSelectionMergeResponse) GetBody() *RunTopicSelectionMergeResponseBody {
	return s.Body
}

func (s *RunTopicSelectionMergeResponse) SetHeaders(v map[string]*string) *RunTopicSelectionMergeResponse {
	s.Headers = v
	return s
}

func (s *RunTopicSelectionMergeResponse) SetStatusCode(v int32) *RunTopicSelectionMergeResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTopicSelectionMergeResponse) SetBody(v *RunTopicSelectionMergeResponseBody) *RunTopicSelectionMergeResponse {
	s.Body = v
	return s
}

func (s *RunTopicSelectionMergeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
