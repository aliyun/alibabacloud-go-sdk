// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocStructureResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocStructureResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocStructureResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDocStructureResultResponseBody) *GetDocStructureResultResponse
	GetBody() *GetDocStructureResultResponseBody
}

type GetDocStructureResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocStructureResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocStructureResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocStructureResultResponse) GoString() string {
	return s.String()
}

func (s *GetDocStructureResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocStructureResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocStructureResultResponse) GetBody() *GetDocStructureResultResponseBody {
	return s.Body
}

func (s *GetDocStructureResultResponse) SetHeaders(v map[string]*string) *GetDocStructureResultResponse {
	s.Headers = v
	return s
}

func (s *GetDocStructureResultResponse) SetStatusCode(v int32) *GetDocStructureResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocStructureResultResponse) SetBody(v *GetDocStructureResultResponseBody) *GetDocStructureResultResponse {
	s.Body = v
	return s
}

func (s *GetDocStructureResultResponse) Validate() error {
	return dara.Validate(s)
}
