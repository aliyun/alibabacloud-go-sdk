// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRTCStatsV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveRTCStatsV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveRTCStatsV2Response
	GetStatusCode() *int32
	SetBody(v *SaveRTCStatsV2ResponseBody) *SaveRTCStatsV2Response
	GetBody() *SaveRTCStatsV2ResponseBody
}

type SaveRTCStatsV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveRTCStatsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveRTCStatsV2Response) String() string {
	return dara.Prettify(s)
}

func (s SaveRTCStatsV2Response) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveRTCStatsV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveRTCStatsV2Response) GetBody() *SaveRTCStatsV2ResponseBody {
	return s.Body
}

func (s *SaveRTCStatsV2Response) SetHeaders(v map[string]*string) *SaveRTCStatsV2Response {
	s.Headers = v
	return s
}

func (s *SaveRTCStatsV2Response) SetStatusCode(v int32) *SaveRTCStatsV2Response {
	s.StatusCode = &v
	return s
}

func (s *SaveRTCStatsV2Response) SetBody(v *SaveRTCStatsV2ResponseBody) *SaveRTCStatsV2Response {
	s.Body = v
	return s
}

func (s *SaveRTCStatsV2Response) Validate() error {
	return dara.Validate(s)
}
