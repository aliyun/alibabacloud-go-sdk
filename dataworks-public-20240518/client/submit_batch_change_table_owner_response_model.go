// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchChangeTableOwnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitBatchChangeTableOwnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitBatchChangeTableOwnerResponse
	GetStatusCode() *int32
	SetBody(v *SubmitBatchChangeTableOwnerResponseBody) *SubmitBatchChangeTableOwnerResponse
	GetBody() *SubmitBatchChangeTableOwnerResponseBody
}

type SubmitBatchChangeTableOwnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitBatchChangeTableOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitBatchChangeTableOwnerResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchChangeTableOwnerResponse) GoString() string {
	return s.String()
}

func (s *SubmitBatchChangeTableOwnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitBatchChangeTableOwnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitBatchChangeTableOwnerResponse) GetBody() *SubmitBatchChangeTableOwnerResponseBody {
	return s.Body
}

func (s *SubmitBatchChangeTableOwnerResponse) SetHeaders(v map[string]*string) *SubmitBatchChangeTableOwnerResponse {
	s.Headers = v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponse) SetStatusCode(v int32) *SubmitBatchChangeTableOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponse) SetBody(v *SubmitBatchChangeTableOwnerResponseBody) *SubmitBatchChangeTableOwnerResponse {
	s.Body = v
	return s
}

func (s *SubmitBatchChangeTableOwnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
