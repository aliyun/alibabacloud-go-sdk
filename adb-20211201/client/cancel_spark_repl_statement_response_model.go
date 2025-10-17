// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSparkReplStatementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSparkReplStatementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSparkReplStatementResponse
	GetStatusCode() *int32
	SetBody(v *CancelSparkReplStatementResponseBody) *CancelSparkReplStatementResponse
	GetBody() *CancelSparkReplStatementResponseBody
}

type CancelSparkReplStatementResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSparkReplStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSparkReplStatementResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSparkReplStatementResponse) GoString() string {
	return s.String()
}

func (s *CancelSparkReplStatementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSparkReplStatementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSparkReplStatementResponse) GetBody() *CancelSparkReplStatementResponseBody {
	return s.Body
}

func (s *CancelSparkReplStatementResponse) SetHeaders(v map[string]*string) *CancelSparkReplStatementResponse {
	s.Headers = v
	return s
}

func (s *CancelSparkReplStatementResponse) SetStatusCode(v int32) *CancelSparkReplStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSparkReplStatementResponse) SetBody(v *CancelSparkReplStatementResponseBody) *CancelSparkReplStatementResponse {
	s.Body = v
	return s
}

func (s *CancelSparkReplStatementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
