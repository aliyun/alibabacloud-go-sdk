// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAICoachScriptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAICoachScriptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAICoachScriptResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAICoachScriptResponseBody) *DeleteAICoachScriptResponse
	GetBody() *DeleteAICoachScriptResponseBody
}

type DeleteAICoachScriptResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAICoachScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAICoachScriptResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAICoachScriptResponse) GoString() string {
	return s.String()
}

func (s *DeleteAICoachScriptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAICoachScriptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAICoachScriptResponse) GetBody() *DeleteAICoachScriptResponseBody {
	return s.Body
}

func (s *DeleteAICoachScriptResponse) SetHeaders(v map[string]*string) *DeleteAICoachScriptResponse {
	s.Headers = v
	return s
}

func (s *DeleteAICoachScriptResponse) SetStatusCode(v int32) *DeleteAICoachScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAICoachScriptResponse) SetBody(v *DeleteAICoachScriptResponseBody) *DeleteAICoachScriptResponse {
	s.Body = v
	return s
}

func (s *DeleteAICoachScriptResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
