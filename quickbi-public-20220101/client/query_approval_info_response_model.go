// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApprovalInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryApprovalInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryApprovalInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryApprovalInfoResponseBody) *QueryApprovalInfoResponse
	GetBody() *QueryApprovalInfoResponseBody
}

type QueryApprovalInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryApprovalInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryApprovalInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryApprovalInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryApprovalInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryApprovalInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryApprovalInfoResponse) GetBody() *QueryApprovalInfoResponseBody {
	return s.Body
}

func (s *QueryApprovalInfoResponse) SetHeaders(v map[string]*string) *QueryApprovalInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryApprovalInfoResponse) SetStatusCode(v int32) *QueryApprovalInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryApprovalInfoResponse) SetBody(v *QueryApprovalInfoResponseBody) *QueryApprovalInfoResponse {
	s.Body = v
	return s
}

func (s *QueryApprovalInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
