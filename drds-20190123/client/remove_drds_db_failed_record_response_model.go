// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDrdsDbFailedRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveDrdsDbFailedRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveDrdsDbFailedRecordResponse
	GetStatusCode() *int32
	SetBody(v *RemoveDrdsDbFailedRecordResponseBody) *RemoveDrdsDbFailedRecordResponse
	GetBody() *RemoveDrdsDbFailedRecordResponseBody
}

type RemoveDrdsDbFailedRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDrdsDbFailedRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDrdsDbFailedRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveDrdsDbFailedRecordResponse) GoString() string {
	return s.String()
}

func (s *RemoveDrdsDbFailedRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveDrdsDbFailedRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveDrdsDbFailedRecordResponse) GetBody() *RemoveDrdsDbFailedRecordResponseBody {
	return s.Body
}

func (s *RemoveDrdsDbFailedRecordResponse) SetHeaders(v map[string]*string) *RemoveDrdsDbFailedRecordResponse {
	s.Headers = v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponse) SetStatusCode(v int32) *RemoveDrdsDbFailedRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponse) SetBody(v *RemoveDrdsDbFailedRecordResponseBody) *RemoveDrdsDbFailedRecordResponse {
	s.Body = v
	return s
}

func (s *RemoveDrdsDbFailedRecordResponse) Validate() error {
	return dara.Validate(s)
}
