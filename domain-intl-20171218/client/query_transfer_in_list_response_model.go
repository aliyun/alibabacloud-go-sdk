// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferInListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTransferInListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTransferInListResponse
	GetStatusCode() *int32
	SetBody(v *QueryTransferInListResponseBody) *QueryTransferInListResponse
	GetBody() *QueryTransferInListResponseBody
}

type QueryTransferInListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTransferInListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTransferInListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInListResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferInListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTransferInListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTransferInListResponse) GetBody() *QueryTransferInListResponseBody {
	return s.Body
}

func (s *QueryTransferInListResponse) SetHeaders(v map[string]*string) *QueryTransferInListResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferInListResponse) SetStatusCode(v int32) *QueryTransferInListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferInListResponse) SetBody(v *QueryTransferInListResponseBody) *QueryTransferInListResponse {
	s.Body = v
	return s
}

func (s *QueryTransferInListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
