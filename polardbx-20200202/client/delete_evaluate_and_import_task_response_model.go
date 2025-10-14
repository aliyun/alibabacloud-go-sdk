// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluateAndImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEvaluateAndImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEvaluateAndImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEvaluateAndImportTaskResponseBody) *DeleteEvaluateAndImportTaskResponse
	GetBody() *DeleteEvaluateAndImportTaskResponseBody
}

type DeleteEvaluateAndImportTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEvaluateAndImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEvaluateAndImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluateAndImportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteEvaluateAndImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEvaluateAndImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEvaluateAndImportTaskResponse) GetBody() *DeleteEvaluateAndImportTaskResponseBody {
	return s.Body
}

func (s *DeleteEvaluateAndImportTaskResponse) SetHeaders(v map[string]*string) *DeleteEvaluateAndImportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteEvaluateAndImportTaskResponse) SetStatusCode(v int32) *DeleteEvaluateAndImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEvaluateAndImportTaskResponse) SetBody(v *DeleteEvaluateAndImportTaskResponseBody) *DeleteEvaluateAndImportTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteEvaluateAndImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
