// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachCheatDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachCheatDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachCheatDetectionResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachCheatDetectionResponseBody) *GetAICoachCheatDetectionResponse
	GetBody() *GetAICoachCheatDetectionResponseBody
}

type GetAICoachCheatDetectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachCheatDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachCheatDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachCheatDetectionResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachCheatDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachCheatDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachCheatDetectionResponse) GetBody() *GetAICoachCheatDetectionResponseBody {
	return s.Body
}

func (s *GetAICoachCheatDetectionResponse) SetHeaders(v map[string]*string) *GetAICoachCheatDetectionResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachCheatDetectionResponse) SetStatusCode(v int32) *GetAICoachCheatDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachCheatDetectionResponse) SetBody(v *GetAICoachCheatDetectionResponseBody) *GetAICoachCheatDetectionResponse {
	s.Body = v
	return s
}

func (s *GetAICoachCheatDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
