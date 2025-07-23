// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResource4ModifyRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResource4ModifyRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResource4ModifyRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetResource4ModifyRecordResponseBody) *GetResource4ModifyRecordResponse
	GetBody() *GetResource4ModifyRecordResponseBody
}

type GetResource4ModifyRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResource4ModifyRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResource4ModifyRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResource4ModifyRecordResponse) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResource4ModifyRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResource4ModifyRecordResponse) GetBody() *GetResource4ModifyRecordResponseBody {
	return s.Body
}

func (s *GetResource4ModifyRecordResponse) SetHeaders(v map[string]*string) *GetResource4ModifyRecordResponse {
	s.Headers = v
	return s
}

func (s *GetResource4ModifyRecordResponse) SetStatusCode(v int32) *GetResource4ModifyRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResource4ModifyRecordResponse) SetBody(v *GetResource4ModifyRecordResponseBody) *GetResource4ModifyRecordResponse {
	s.Body = v
	return s
}

func (s *GetResource4ModifyRecordResponse) Validate() error {
	return dara.Validate(s)
}
