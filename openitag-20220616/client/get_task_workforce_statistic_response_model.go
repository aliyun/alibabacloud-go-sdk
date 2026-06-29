// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskWorkforceStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTaskWorkforceStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTaskWorkforceStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetTaskWorkforceStatisticResponseBody) *GetTaskWorkforceStatisticResponse
	GetBody() *GetTaskWorkforceStatisticResponseBody
}

type GetTaskWorkforceStatisticResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskWorkforceStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskWorkforceStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTaskWorkforceStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTaskWorkforceStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTaskWorkforceStatisticResponse) GetBody() *GetTaskWorkforceStatisticResponseBody {
	return s.Body
}

func (s *GetTaskWorkforceStatisticResponse) SetHeaders(v map[string]*string) *GetTaskWorkforceStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetTaskWorkforceStatisticResponse) SetStatusCode(v int32) *GetTaskWorkforceStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponse) SetBody(v *GetTaskWorkforceStatisticResponseBody) *GetTaskWorkforceStatisticResponse {
	s.Body = v
	return s
}

func (s *GetTaskWorkforceStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
