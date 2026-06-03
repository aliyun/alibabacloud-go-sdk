// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderActivateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderActivateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderActivateResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForCreatingOrderActivateResponseBody) *SaveBatchTaskForCreatingOrderActivateResponse
	GetBody() *SaveBatchTaskForCreatingOrderActivateResponseBody
}

type SaveBatchTaskForCreatingOrderActivateResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForCreatingOrderActivateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderActivateResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderActivateResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) GetBody() *SaveBatchTaskForCreatingOrderActivateResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderActivateResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderActivateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) SetBody(v *SaveBatchTaskForCreatingOrderActivateResponseBody) *SaveBatchTaskForCreatingOrderActivateResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderActivateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
