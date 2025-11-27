// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPurchaseControlRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPurchaseControlRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPurchaseControlRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetPurchaseControlRecordResponseBody) *GetPurchaseControlRecordResponse
	GetBody() *GetPurchaseControlRecordResponseBody
}

type GetPurchaseControlRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPurchaseControlRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPurchaseControlRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPurchaseControlRecordResponse) GoString() string {
	return s.String()
}

func (s *GetPurchaseControlRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPurchaseControlRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPurchaseControlRecordResponse) GetBody() *GetPurchaseControlRecordResponseBody {
	return s.Body
}

func (s *GetPurchaseControlRecordResponse) SetHeaders(v map[string]*string) *GetPurchaseControlRecordResponse {
	s.Headers = v
	return s
}

func (s *GetPurchaseControlRecordResponse) SetStatusCode(v int32) *GetPurchaseControlRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPurchaseControlRecordResponse) SetBody(v *GetPurchaseControlRecordResponseBody) *GetPurchaseControlRecordResponse {
	s.Body = v
	return s
}

func (s *GetPurchaseControlRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
