// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterveneImportTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInterveneImportTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInterveneImportTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetInterveneImportTaskInfoResponseBody) *GetInterveneImportTaskInfoResponse
	GetBody() *GetInterveneImportTaskInfoResponseBody
}

type GetInterveneImportTaskInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneImportTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneImportTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInterveneImportTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInterveneImportTaskInfoResponse) GetBody() *GetInterveneImportTaskInfoResponseBody {
	return s.Body
}

func (s *GetInterveneImportTaskInfoResponse) SetHeaders(v map[string]*string) *GetInterveneImportTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneImportTaskInfoResponse) SetStatusCode(v int32) *GetInterveneImportTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponse) SetBody(v *GetInterveneImportTaskInfoResponseBody) *GetInterveneImportTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetInterveneImportTaskInfoResponse) Validate() error {
	return dara.Validate(s)
}
