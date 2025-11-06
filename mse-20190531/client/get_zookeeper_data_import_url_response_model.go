// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetZookeeperDataImportUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetZookeeperDataImportUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetZookeeperDataImportUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetZookeeperDataImportUrlResponseBody) *GetZookeeperDataImportUrlResponse
	GetBody() *GetZookeeperDataImportUrlResponseBody
}

type GetZookeeperDataImportUrlResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetZookeeperDataImportUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetZookeeperDataImportUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetZookeeperDataImportUrlResponse) GoString() string {
	return s.String()
}

func (s *GetZookeeperDataImportUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetZookeeperDataImportUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetZookeeperDataImportUrlResponse) GetBody() *GetZookeeperDataImportUrlResponseBody {
	return s.Body
}

func (s *GetZookeeperDataImportUrlResponse) SetHeaders(v map[string]*string) *GetZookeeperDataImportUrlResponse {
	s.Headers = v
	return s
}

func (s *GetZookeeperDataImportUrlResponse) SetStatusCode(v int32) *GetZookeeperDataImportUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetZookeeperDataImportUrlResponse) SetBody(v *GetZookeeperDataImportUrlResponseBody) *GetZookeeperDataImportUrlResponse {
	s.Body = v
	return s
}

func (s *GetZookeeperDataImportUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
