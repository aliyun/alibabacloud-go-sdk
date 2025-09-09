// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanOssObjectV1Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScanOssObjectV1Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScanOssObjectV1Response
	GetStatusCode() *int32
	SetBody(v *ScanOssObjectV1ResponseBody) *ScanOssObjectV1Response
	GetBody() *ScanOssObjectV1ResponseBody
}

type ScanOssObjectV1Response struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanOssObjectV1ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanOssObjectV1Response) String() string {
	return dara.Prettify(s)
}

func (s ScanOssObjectV1Response) GoString() string {
	return s.String()
}

func (s *ScanOssObjectV1Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScanOssObjectV1Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScanOssObjectV1Response) GetBody() *ScanOssObjectV1ResponseBody {
	return s.Body
}

func (s *ScanOssObjectV1Response) SetHeaders(v map[string]*string) *ScanOssObjectV1Response {
	s.Headers = v
	return s
}

func (s *ScanOssObjectV1Response) SetStatusCode(v int32) *ScanOssObjectV1Response {
	s.StatusCode = &v
	return s
}

func (s *ScanOssObjectV1Response) SetBody(v *ScanOssObjectV1ResponseBody) *ScanOssObjectV1Response {
	s.Body = v
	return s
}

func (s *ScanOssObjectV1Response) Validate() error {
	return dara.Validate(s)
}
