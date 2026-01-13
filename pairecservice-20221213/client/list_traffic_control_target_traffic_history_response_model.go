// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficControlTargetTrafficHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrafficControlTargetTrafficHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrafficControlTargetTrafficHistoryResponse
	GetStatusCode() *int32
	SetBody(v *ListTrafficControlTargetTrafficHistoryResponseBody) *ListTrafficControlTargetTrafficHistoryResponse
	GetBody() *ListTrafficControlTargetTrafficHistoryResponseBody
}

type ListTrafficControlTargetTrafficHistoryResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficControlTargetTrafficHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficControlTargetTrafficHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTargetTrafficHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) GetBody() *ListTrafficControlTargetTrafficHistoryResponseBody {
	return s.Body
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) SetHeaders(v map[string]*string) *ListTrafficControlTargetTrafficHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) SetStatusCode(v int32) *ListTrafficControlTargetTrafficHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) SetBody(v *ListTrafficControlTargetTrafficHistoryResponseBody) *ListTrafficControlTargetTrafficHistoryResponse {
	s.Body = v
	return s
}

func (s *ListTrafficControlTargetTrafficHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
