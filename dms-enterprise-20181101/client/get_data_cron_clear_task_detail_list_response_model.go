// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCronClearTaskDetailListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCronClearTaskDetailListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCronClearTaskDetailListResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCronClearTaskDetailListResponseBody) *GetDataCronClearTaskDetailListResponse
	GetBody() *GetDataCronClearTaskDetailListResponseBody
}

type GetDataCronClearTaskDetailListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCronClearTaskDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCronClearTaskDetailListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearTaskDetailListResponse) GoString() string {
	return s.String()
}

func (s *GetDataCronClearTaskDetailListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCronClearTaskDetailListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCronClearTaskDetailListResponse) GetBody() *GetDataCronClearTaskDetailListResponseBody {
	return s.Body
}

func (s *GetDataCronClearTaskDetailListResponse) SetHeaders(v map[string]*string) *GetDataCronClearTaskDetailListResponse {
	s.Headers = v
	return s
}

func (s *GetDataCronClearTaskDetailListResponse) SetStatusCode(v int32) *GetDataCronClearTaskDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCronClearTaskDetailListResponse) SetBody(v *GetDataCronClearTaskDetailListResponseBody) *GetDataCronClearTaskDetailListResponse {
	s.Body = v
	return s
}

func (s *GetDataCronClearTaskDetailListResponse) Validate() error {
	return dara.Validate(s)
}
