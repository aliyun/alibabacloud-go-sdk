// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallDetailRecordsV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallDetailRecordsV2Response
	GetStatusCode() *int32
	SetBody(v *ListCallDetailRecordsV2ResponseBody) *ListCallDetailRecordsV2Response
	GetBody() *ListCallDetailRecordsV2ResponseBody
}

type ListCallDetailRecordsV2Response struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallDetailRecordsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallDetailRecordsV2Response) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsV2Response) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallDetailRecordsV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallDetailRecordsV2Response) GetBody() *ListCallDetailRecordsV2ResponseBody {
	return s.Body
}

func (s *ListCallDetailRecordsV2Response) SetHeaders(v map[string]*string) *ListCallDetailRecordsV2Response {
	s.Headers = v
	return s
}

func (s *ListCallDetailRecordsV2Response) SetStatusCode(v int32) *ListCallDetailRecordsV2Response {
	s.StatusCode = &v
	return s
}

func (s *ListCallDetailRecordsV2Response) SetBody(v *ListCallDetailRecordsV2ResponseBody) *ListCallDetailRecordsV2Response {
	s.Body = v
	return s
}

func (s *ListCallDetailRecordsV2Response) Validate() error {
	return dara.Validate(s)
}
