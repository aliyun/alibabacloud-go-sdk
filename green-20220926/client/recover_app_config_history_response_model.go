// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAppConfigHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverAppConfigHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverAppConfigHistoryResponse
	GetStatusCode() *int32
	SetBody(v *RecoverAppConfigHistoryResponseBody) *RecoverAppConfigHistoryResponse
	GetBody() *RecoverAppConfigHistoryResponseBody
}

type RecoverAppConfigHistoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverAppConfigHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverAppConfigHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverAppConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *RecoverAppConfigHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverAppConfigHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverAppConfigHistoryResponse) GetBody() *RecoverAppConfigHistoryResponseBody {
	return s.Body
}

func (s *RecoverAppConfigHistoryResponse) SetHeaders(v map[string]*string) *RecoverAppConfigHistoryResponse {
	s.Headers = v
	return s
}

func (s *RecoverAppConfigHistoryResponse) SetStatusCode(v int32) *RecoverAppConfigHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverAppConfigHistoryResponse) SetBody(v *RecoverAppConfigHistoryResponseBody) *RecoverAppConfigHistoryResponse {
	s.Body = v
	return s
}

func (s *RecoverAppConfigHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
