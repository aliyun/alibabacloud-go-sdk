// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineWaitingNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineWaitingNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineWaitingNumberResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineWaitingNumberResponseBody) *GetHotlineWaitingNumberResponse
	GetBody() *GetHotlineWaitingNumberResponseBody
}

type GetHotlineWaitingNumberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineWaitingNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineWaitingNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineWaitingNumberResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineWaitingNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineWaitingNumberResponse) GetBody() *GetHotlineWaitingNumberResponseBody {
	return s.Body
}

func (s *GetHotlineWaitingNumberResponse) SetHeaders(v map[string]*string) *GetHotlineWaitingNumberResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineWaitingNumberResponse) SetStatusCode(v int32) *GetHotlineWaitingNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineWaitingNumberResponse) SetBody(v *GetHotlineWaitingNumberResponseBody) *GetHotlineWaitingNumberResponse {
	s.Body = v
	return s
}

func (s *GetHotlineWaitingNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
