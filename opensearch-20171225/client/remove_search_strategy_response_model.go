// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSearchStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveSearchStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveSearchStrategyResponse
	GetStatusCode() *int32
	SetBody(v *RemoveSearchStrategyResponseBody) *RemoveSearchStrategyResponse
	GetBody() *RemoveSearchStrategyResponseBody
}

type RemoveSearchStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveSearchStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *RemoveSearchStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveSearchStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveSearchStrategyResponse) GetBody() *RemoveSearchStrategyResponseBody {
	return s.Body
}

func (s *RemoveSearchStrategyResponse) SetHeaders(v map[string]*string) *RemoveSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *RemoveSearchStrategyResponse) SetStatusCode(v int32) *RemoveSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSearchStrategyResponse) SetBody(v *RemoveSearchStrategyResponseBody) *RemoveSearchStrategyResponse {
	s.Body = v
	return s
}

func (s *RemoveSearchStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
