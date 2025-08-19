// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageContactDeletionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageContactDeletionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageContactDeletionStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageContactDeletionStatusResponseBody) *GetMessageContactDeletionStatusResponse
	GetBody() *GetMessageContactDeletionStatusResponseBody
}

type GetMessageContactDeletionStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageContactDeletionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageContactDeletionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageContactDeletionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMessageContactDeletionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageContactDeletionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageContactDeletionStatusResponse) GetBody() *GetMessageContactDeletionStatusResponseBody {
	return s.Body
}

func (s *GetMessageContactDeletionStatusResponse) SetHeaders(v map[string]*string) *GetMessageContactDeletionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMessageContactDeletionStatusResponse) SetStatusCode(v int32) *GetMessageContactDeletionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageContactDeletionStatusResponse) SetBody(v *GetMessageContactDeletionStatusResponseBody) *GetMessageContactDeletionStatusResponse {
	s.Body = v
	return s
}

func (s *GetMessageContactDeletionStatusResponse) Validate() error {
	return dara.Validate(s)
}
