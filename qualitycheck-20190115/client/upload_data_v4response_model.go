// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadDataV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadDataV4Response
	GetStatusCode() *int32
	SetBody(v *UploadDataV4ResponseBody) *UploadDataV4Response
	GetBody() *UploadDataV4ResponseBody
}

type UploadDataV4Response struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDataV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDataV4Response) String() string {
	return dara.Prettify(s)
}

func (s UploadDataV4Response) GoString() string {
	return s.String()
}

func (s *UploadDataV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadDataV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadDataV4Response) GetBody() *UploadDataV4ResponseBody {
	return s.Body
}

func (s *UploadDataV4Response) SetHeaders(v map[string]*string) *UploadDataV4Response {
	s.Headers = v
	return s
}

func (s *UploadDataV4Response) SetStatusCode(v int32) *UploadDataV4Response {
	s.StatusCode = &v
	return s
}

func (s *UploadDataV4Response) SetBody(v *UploadDataV4ResponseBody) *UploadDataV4Response {
	s.Body = v
	return s
}

func (s *UploadDataV4Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
