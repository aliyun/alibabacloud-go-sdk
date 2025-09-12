// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormFsUsedDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLindormFsUsedDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLindormFsUsedDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetLindormFsUsedDetailResponseBody) *GetLindormFsUsedDetailResponse
	GetBody() *GetLindormFsUsedDetailResponseBody
}

type GetLindormFsUsedDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormFsUsedDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormFsUsedDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLindormFsUsedDetailResponse) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLindormFsUsedDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLindormFsUsedDetailResponse) GetBody() *GetLindormFsUsedDetailResponseBody {
	return s.Body
}

func (s *GetLindormFsUsedDetailResponse) SetHeaders(v map[string]*string) *GetLindormFsUsedDetailResponse {
	s.Headers = v
	return s
}

func (s *GetLindormFsUsedDetailResponse) SetStatusCode(v int32) *GetLindormFsUsedDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormFsUsedDetailResponse) SetBody(v *GetLindormFsUsedDetailResponseBody) *GetLindormFsUsedDetailResponse {
	s.Body = v
	return s
}

func (s *GetLindormFsUsedDetailResponse) Validate() error {
	return dara.Validate(s)
}
