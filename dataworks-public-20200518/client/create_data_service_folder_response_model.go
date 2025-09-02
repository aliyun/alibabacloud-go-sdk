// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceFolderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataServiceFolderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataServiceFolderResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataServiceFolderResponseBody) *CreateDataServiceFolderResponse
	GetBody() *CreateDataServiceFolderResponseBody
}

type CreateDataServiceFolderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataServiceFolderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataServiceFolderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceFolderResponse) GoString() string {
	return s.String()
}

func (s *CreateDataServiceFolderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataServiceFolderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataServiceFolderResponse) GetBody() *CreateDataServiceFolderResponseBody {
	return s.Body
}

func (s *CreateDataServiceFolderResponse) SetHeaders(v map[string]*string) *CreateDataServiceFolderResponse {
	s.Headers = v
	return s
}

func (s *CreateDataServiceFolderResponse) SetStatusCode(v int32) *CreateDataServiceFolderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataServiceFolderResponse) SetBody(v *CreateDataServiceFolderResponseBody) *CreateDataServiceFolderResponse {
	s.Body = v
	return s
}

func (s *CreateDataServiceFolderResponse) Validate() error {
	return dara.Validate(s)
}
