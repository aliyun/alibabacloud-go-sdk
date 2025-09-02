// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableFullInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaTableFullInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaTableFullInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaTableFullInfoResponseBody) *GetMetaTableFullInfoResponse
	GetBody() *GetMetaTableFullInfoResponseBody
}

type GetMetaTableFullInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaTableFullInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaTableFullInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableFullInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMetaTableFullInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaTableFullInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaTableFullInfoResponse) GetBody() *GetMetaTableFullInfoResponseBody {
	return s.Body
}

func (s *GetMetaTableFullInfoResponse) SetHeaders(v map[string]*string) *GetMetaTableFullInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMetaTableFullInfoResponse) SetStatusCode(v int32) *GetMetaTableFullInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaTableFullInfoResponse) SetBody(v *GetMetaTableFullInfoResponseBody) *GetMetaTableFullInfoResponse {
	s.Body = v
	return s
}

func (s *GetMetaTableFullInfoResponse) Validate() error {
	return dara.Validate(s)
}
