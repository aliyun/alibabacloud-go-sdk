// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallInPoolTransferConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCallInPoolTransferConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCallInPoolTransferConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryCallInPoolTransferConfigResponseBody) *QueryCallInPoolTransferConfigResponse
	GetBody() *QueryCallInPoolTransferConfigResponseBody
}

type QueryCallInPoolTransferConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallInPoolTransferConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCallInPoolTransferConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCallInPoolTransferConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCallInPoolTransferConfigResponse) GetBody() *QueryCallInPoolTransferConfigResponseBody {
	return s.Body
}

func (s *QueryCallInPoolTransferConfigResponse) SetHeaders(v map[string]*string) *QueryCallInPoolTransferConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryCallInPoolTransferConfigResponse) SetStatusCode(v int32) *QueryCallInPoolTransferConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponse) SetBody(v *QueryCallInPoolTransferConfigResponseBody) *QueryCallInPoolTransferConfigResponse {
	s.Body = v
	return s
}

func (s *QueryCallInPoolTransferConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
