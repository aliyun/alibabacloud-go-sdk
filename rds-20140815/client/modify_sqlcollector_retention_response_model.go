// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLCollectorRetentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySQLCollectorRetentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySQLCollectorRetentionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySQLCollectorRetentionResponseBody) *ModifySQLCollectorRetentionResponse
	GetBody() *ModifySQLCollectorRetentionResponseBody
}

type ModifySQLCollectorRetentionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySQLCollectorRetentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySQLCollectorRetentionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLCollectorRetentionResponse) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorRetentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySQLCollectorRetentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySQLCollectorRetentionResponse) GetBody() *ModifySQLCollectorRetentionResponseBody {
	return s.Body
}

func (s *ModifySQLCollectorRetentionResponse) SetHeaders(v map[string]*string) *ModifySQLCollectorRetentionResponse {
	s.Headers = v
	return s
}

func (s *ModifySQLCollectorRetentionResponse) SetStatusCode(v int32) *ModifySQLCollectorRetentionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySQLCollectorRetentionResponse) SetBody(v *ModifySQLCollectorRetentionResponseBody) *ModifySQLCollectorRetentionResponse {
	s.Body = v
	return s
}

func (s *ModifySQLCollectorRetentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
