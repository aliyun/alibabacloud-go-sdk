// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFullRequestStatResultByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFullRequestStatResultByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFullRequestStatResultByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *GetFullRequestStatResultByInstanceIdResponseBody) *GetFullRequestStatResultByInstanceIdResponse
	GetBody() *GetFullRequestStatResultByInstanceIdResponseBody
}

type GetFullRequestStatResultByInstanceIdResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFullRequestStatResultByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFullRequestStatResultByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFullRequestStatResultByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *GetFullRequestStatResultByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFullRequestStatResultByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFullRequestStatResultByInstanceIdResponse) GetBody() *GetFullRequestStatResultByInstanceIdResponseBody {
	return s.Body
}

func (s *GetFullRequestStatResultByInstanceIdResponse) SetHeaders(v map[string]*string) *GetFullRequestStatResultByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponse) SetStatusCode(v int32) *GetFullRequestStatResultByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponse) SetBody(v *GetFullRequestStatResultByInstanceIdResponseBody) *GetFullRequestStatResultByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *GetFullRequestStatResultByInstanceIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
