// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceStoreInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceStoreInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceStoreInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceStoreInfoResponseBody) *GetInstanceStoreInfoResponse
	GetBody() *GetInstanceStoreInfoResponseBody
}

type GetInstanceStoreInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceStoreInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceStoreInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceStoreInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceStoreInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceStoreInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceStoreInfoResponse) GetBody() *GetInstanceStoreInfoResponseBody {
	return s.Body
}

func (s *GetInstanceStoreInfoResponse) SetHeaders(v map[string]*string) *GetInstanceStoreInfoResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceStoreInfoResponse) SetStatusCode(v int32) *GetInstanceStoreInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceStoreInfoResponse) SetBody(v *GetInstanceStoreInfoResponseBody) *GetInstanceStoreInfoResponse {
	s.Body = v
	return s
}

func (s *GetInstanceStoreInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
