// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetContentJobIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSheetContentJobIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSheetContentJobIdResponse
	GetStatusCode() *int32
	SetBody(v *GetSheetContentJobIdResponseBody) *GetSheetContentJobIdResponse
	GetBody() *GetSheetContentJobIdResponseBody
}

type GetSheetContentJobIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSheetContentJobIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSheetContentJobIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdResponse) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSheetContentJobIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSheetContentJobIdResponse) GetBody() *GetSheetContentJobIdResponseBody {
	return s.Body
}

func (s *GetSheetContentJobIdResponse) SetHeaders(v map[string]*string) *GetSheetContentJobIdResponse {
	s.Headers = v
	return s
}

func (s *GetSheetContentJobIdResponse) SetStatusCode(v int32) *GetSheetContentJobIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSheetContentJobIdResponse) SetBody(v *GetSheetContentJobIdResponseBody) *GetSheetContentJobIdResponse {
	s.Body = v
	return s
}

func (s *GetSheetContentJobIdResponse) Validate() error {
	return dara.Validate(s)
}
