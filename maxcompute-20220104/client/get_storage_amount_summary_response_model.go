// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageAmountSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageAmountSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageAmountSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageAmountSummaryResponseBody) *GetStorageAmountSummaryResponse
	GetBody() *GetStorageAmountSummaryResponseBody
}

type GetStorageAmountSummaryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageAmountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageAmountSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageAmountSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetStorageAmountSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageAmountSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageAmountSummaryResponse) GetBody() *GetStorageAmountSummaryResponseBody {
	return s.Body
}

func (s *GetStorageAmountSummaryResponse) SetHeaders(v map[string]*string) *GetStorageAmountSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetStorageAmountSummaryResponse) SetStatusCode(v int32) *GetStorageAmountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageAmountSummaryResponse) SetBody(v *GetStorageAmountSummaryResponseBody) *GetStorageAmountSummaryResponse {
	s.Body = v
	return s
}

func (s *GetStorageAmountSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
