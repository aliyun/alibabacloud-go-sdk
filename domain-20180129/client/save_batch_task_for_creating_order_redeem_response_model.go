// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForCreatingOrderRedeemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderRedeemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderRedeemResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForCreatingOrderRedeemResponseBody) *SaveBatchTaskForCreatingOrderRedeemResponse
	GetBody() *SaveBatchTaskForCreatingOrderRedeemResponseBody
}

type SaveBatchTaskForCreatingOrderRedeemResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForCreatingOrderRedeemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForCreatingOrderRedeemResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForCreatingOrderRedeemResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) GetBody() *SaveBatchTaskForCreatingOrderRedeemResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForCreatingOrderRedeemResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) SetStatusCode(v int32) *SaveBatchTaskForCreatingOrderRedeemResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) SetBody(v *SaveBatchTaskForCreatingOrderRedeemResponseBody) *SaveBatchTaskForCreatingOrderRedeemResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForCreatingOrderRedeemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
