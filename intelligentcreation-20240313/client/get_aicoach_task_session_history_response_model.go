// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICoachTaskSessionHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICoachTaskSessionHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICoachTaskSessionHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetAICoachTaskSessionHistoryResponseBody) *GetAICoachTaskSessionHistoryResponse
	GetBody() *GetAICoachTaskSessionHistoryResponseBody
}

type GetAICoachTaskSessionHistoryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICoachTaskSessionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICoachTaskSessionHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICoachTaskSessionHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetAICoachTaskSessionHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICoachTaskSessionHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICoachTaskSessionHistoryResponse) GetBody() *GetAICoachTaskSessionHistoryResponseBody {
	return s.Body
}

func (s *GetAICoachTaskSessionHistoryResponse) SetHeaders(v map[string]*string) *GetAICoachTaskSessionHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponse) SetStatusCode(v int32) *GetAICoachTaskSessionHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponse) SetBody(v *GetAICoachTaskSessionHistoryResponseBody) *GetAICoachTaskSessionHistoryResponse {
	s.Body = v
	return s
}

func (s *GetAICoachTaskSessionHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
