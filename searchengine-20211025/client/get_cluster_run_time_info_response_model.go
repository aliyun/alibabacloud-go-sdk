// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterRunTimeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterRunTimeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterRunTimeInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterRunTimeInfoResponseBody) *GetClusterRunTimeInfoResponse
	GetBody() *GetClusterRunTimeInfoResponseBody
}

type GetClusterRunTimeInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterRunTimeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterRunTimeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRunTimeInfoResponse) GoString() string {
	return s.String()
}

func (s *GetClusterRunTimeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterRunTimeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterRunTimeInfoResponse) GetBody() *GetClusterRunTimeInfoResponseBody {
	return s.Body
}

func (s *GetClusterRunTimeInfoResponse) SetHeaders(v map[string]*string) *GetClusterRunTimeInfoResponse {
	s.Headers = v
	return s
}

func (s *GetClusterRunTimeInfoResponse) SetStatusCode(v int32) *GetClusterRunTimeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterRunTimeInfoResponse) SetBody(v *GetClusterRunTimeInfoResponseBody) *GetClusterRunTimeInfoResponse {
	s.Body = v
	return s
}

func (s *GetClusterRunTimeInfoResponse) Validate() error {
	return dara.Validate(s)
}
