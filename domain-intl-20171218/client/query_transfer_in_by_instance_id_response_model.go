// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferInByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTransferInByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTransferInByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryTransferInByInstanceIdResponseBody) *QueryTransferInByInstanceIdResponse
	GetBody() *QueryTransferInByInstanceIdResponseBody
}

type QueryTransferInByInstanceIdResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTransferInByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTransferInByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferInByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *QueryTransferInByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTransferInByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTransferInByInstanceIdResponse) GetBody() *QueryTransferInByInstanceIdResponseBody {
	return s.Body
}

func (s *QueryTransferInByInstanceIdResponse) SetHeaders(v map[string]*string) *QueryTransferInByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *QueryTransferInByInstanceIdResponse) SetStatusCode(v int32) *QueryTransferInByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTransferInByInstanceIdResponse) SetBody(v *QueryTransferInByInstanceIdResponseBody) *QueryTransferInByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *QueryTransferInByInstanceIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
