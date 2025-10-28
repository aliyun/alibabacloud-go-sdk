// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelStackOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelStackOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelStackOperationResponse
	GetStatusCode() *int32
	SetBody(v *CancelStackOperationResponseBody) *CancelStackOperationResponse
	GetBody() *CancelStackOperationResponseBody
}

type CancelStackOperationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelStackOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelStackOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelStackOperationResponse) GoString() string {
	return s.String()
}

func (s *CancelStackOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelStackOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelStackOperationResponse) GetBody() *CancelStackOperationResponseBody {
	return s.Body
}

func (s *CancelStackOperationResponse) SetHeaders(v map[string]*string) *CancelStackOperationResponse {
	s.Headers = v
	return s
}

func (s *CancelStackOperationResponse) SetStatusCode(v int32) *CancelStackOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelStackOperationResponse) SetBody(v *CancelStackOperationResponseBody) *CancelStackOperationResponse {
	s.Body = v
	return s
}

func (s *CancelStackOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
