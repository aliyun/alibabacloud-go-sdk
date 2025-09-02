// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataServiceFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataServiceFolderResponse
	GetStatusCode() *int32
	SetBody(v *GetDataServiceFolderResponseBody) *GetDataServiceFolderResponse
	GetBody() *GetDataServiceFolderResponseBody
}

type GetDataServiceFolderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataServiceFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataServiceFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceFolderResponse) GoString() string {
	return s.String()
}

func (s *GetDataServiceFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataServiceFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataServiceFolderResponse) GetBody() *GetDataServiceFolderResponseBody {
	return s.Body
}

func (s *GetDataServiceFolderResponse) SetHeaders(v map[string]*string) *GetDataServiceFolderResponse {
	s.Headers = v
	return s
}

func (s *GetDataServiceFolderResponse) SetStatusCode(v int32) *GetDataServiceFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataServiceFolderResponse) SetBody(v *GetDataServiceFolderResponseBody) *GetDataServiceFolderResponse {
	s.Body = v
	return s
}

func (s *GetDataServiceFolderResponse) Validate() error {
	return dara.Validate(s)
}
