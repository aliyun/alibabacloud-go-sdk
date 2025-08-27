// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainExceedApplyQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TrainExceedApplyQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TrainExceedApplyQueryResponse
	GetStatusCode() *int32
	SetBody(v *TrainExceedApplyQueryResponseBody) *TrainExceedApplyQueryResponse
	GetBody() *TrainExceedApplyQueryResponseBody
}

type TrainExceedApplyQueryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrainExceedApplyQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrainExceedApplyQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s TrainExceedApplyQueryResponse) GoString() string {
	return s.String()
}

func (s *TrainExceedApplyQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TrainExceedApplyQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TrainExceedApplyQueryResponse) GetBody() *TrainExceedApplyQueryResponseBody {
	return s.Body
}

func (s *TrainExceedApplyQueryResponse) SetHeaders(v map[string]*string) *TrainExceedApplyQueryResponse {
	s.Headers = v
	return s
}

func (s *TrainExceedApplyQueryResponse) SetStatusCode(v int32) *TrainExceedApplyQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *TrainExceedApplyQueryResponse) SetBody(v *TrainExceedApplyQueryResponseBody) *TrainExceedApplyQueryResponse {
	s.Body = v
	return s
}

func (s *TrainExceedApplyQueryResponse) Validate() error {
	return dara.Validate(s)
}
