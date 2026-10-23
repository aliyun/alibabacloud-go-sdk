// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgenticDBProjectHistoryRetentionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAgenticDBProjectHistoryRetentionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAgenticDBProjectHistoryRetentionResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAgenticDBProjectHistoryRetentionResponseBody) *ModifyAgenticDBProjectHistoryRetentionResponse
	GetBody() *ModifyAgenticDBProjectHistoryRetentionResponseBody
}

type ModifyAgenticDBProjectHistoryRetentionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAgenticDBProjectHistoryRetentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAgenticDBProjectHistoryRetentionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgenticDBProjectHistoryRetentionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponse) GetBody() *ModifyAgenticDBProjectHistoryRetentionResponseBody {
	return s.Body
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponse) SetHeaders(v map[string]*string) *ModifyAgenticDBProjectHistoryRetentionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponse) SetStatusCode(v int32) *ModifyAgenticDBProjectHistoryRetentionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponse) SetBody(v *ModifyAgenticDBProjectHistoryRetentionResponseBody) *ModifyAgenticDBProjectHistoryRetentionResponse {
	s.Body = v
	return s
}

func (s *ModifyAgenticDBProjectHistoryRetentionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
