// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDistributeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageDistributeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageDistributeInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetImageDistributeInfoResponseBody) *GetImageDistributeInfoResponse
	GetBody() *GetImageDistributeInfoResponseBody
}

type GetImageDistributeInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageDistributeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageDistributeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageDistributeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetImageDistributeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageDistributeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageDistributeInfoResponse) GetBody() *GetImageDistributeInfoResponseBody {
	return s.Body
}

func (s *GetImageDistributeInfoResponse) SetHeaders(v map[string]*string) *GetImageDistributeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetImageDistributeInfoResponse) SetStatusCode(v int32) *GetImageDistributeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageDistributeInfoResponse) SetBody(v *GetImageDistributeInfoResponseBody) *GetImageDistributeInfoResponse {
	s.Body = v
	return s
}

func (s *GetImageDistributeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
