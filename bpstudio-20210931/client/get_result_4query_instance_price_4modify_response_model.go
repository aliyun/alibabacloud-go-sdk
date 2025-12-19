// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResult4QueryInstancePrice4ModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResult4QueryInstancePrice4ModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResult4QueryInstancePrice4ModifyResponse
	GetStatusCode() *int32
	SetBody(v *GetResult4QueryInstancePrice4ModifyResponseBody) *GetResult4QueryInstancePrice4ModifyResponse
	GetBody() *GetResult4QueryInstancePrice4ModifyResponseBody
}

type GetResult4QueryInstancePrice4ModifyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResult4QueryInstancePrice4ModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponse) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) GetBody() *GetResult4QueryInstancePrice4ModifyResponseBody {
	return s.Body
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) SetHeaders(v map[string]*string) *GetResult4QueryInstancePrice4ModifyResponse {
	s.Headers = v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) SetStatusCode(v int32) *GetResult4QueryInstancePrice4ModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) SetBody(v *GetResult4QueryInstancePrice4ModifyResponseBody) *GetResult4QueryInstancePrice4ModifyResponse {
	s.Body = v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
