// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferOutInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTransferOutInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTransferOutInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryTransferOutInfoResponseBody) *QueryTransferOutInfoResponse
	GetBody() *QueryTransferOutInfoResponseBody
}

type QueryTransferOutInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTransferOutInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTransferOutInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferOutInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferOutInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTransferOutInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTransferOutInfoResponse) GetBody() *QueryTransferOutInfoResponseBody {
	return s.Body
}

func (s *QueryTransferOutInfoResponse) SetHeaders(v map[string]*string) *QueryTransferOutInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferOutInfoResponse) SetStatusCode(v int32) *QueryTransferOutInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferOutInfoResponse) SetBody(v *QueryTransferOutInfoResponseBody) *QueryTransferOutInfoResponse {
	s.Body = v
	return s
}

func (s *QueryTransferOutInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
