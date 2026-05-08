// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildAICoachScriptRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuildAICoachScriptRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuildAICoachScriptRecordResponse
	GetStatusCode() *int32
	SetBody(v *BuildAICoachScriptRecordResponseBody) *BuildAICoachScriptRecordResponse
	GetBody() *BuildAICoachScriptRecordResponseBody
}

type BuildAICoachScriptRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildAICoachScriptRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildAICoachScriptRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s BuildAICoachScriptRecordResponse) GoString() string {
	return s.String()
}

func (s *BuildAICoachScriptRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuildAICoachScriptRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuildAICoachScriptRecordResponse) GetBody() *BuildAICoachScriptRecordResponseBody {
	return s.Body
}

func (s *BuildAICoachScriptRecordResponse) SetHeaders(v map[string]*string) *BuildAICoachScriptRecordResponse {
	s.Headers = v
	return s
}

func (s *BuildAICoachScriptRecordResponse) SetStatusCode(v int32) *BuildAICoachScriptRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildAICoachScriptRecordResponse) SetBody(v *BuildAICoachScriptRecordResponseBody) *BuildAICoachScriptRecordResponse {
	s.Body = v
	return s
}

func (s *BuildAICoachScriptRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
