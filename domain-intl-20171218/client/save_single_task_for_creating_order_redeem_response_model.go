// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderRedeemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderRedeemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderRedeemResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForCreatingOrderRedeemResponseBody) *SaveSingleTaskForCreatingOrderRedeemResponse
	GetBody() *SaveSingleTaskForCreatingOrderRedeemResponseBody
}

type SaveSingleTaskForCreatingOrderRedeemResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForCreatingOrderRedeemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderRedeemResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderRedeemResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) GetBody() *SaveSingleTaskForCreatingOrderRedeemResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForCreatingOrderRedeemResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) SetStatusCode(v int32) *SaveSingleTaskForCreatingOrderRedeemResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) SetBody(v *SaveSingleTaskForCreatingOrderRedeemResponseBody) *SaveSingleTaskForCreatingOrderRedeemResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForCreatingOrderRedeemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
