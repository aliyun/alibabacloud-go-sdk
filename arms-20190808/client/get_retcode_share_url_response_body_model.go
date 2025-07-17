// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRetcodeShareUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRetcodeShareUrlResponseBody
	GetRequestId() *string
	SetUrl(v string) *GetRetcodeShareUrlResponseBody
	GetUrl() *string
}

type GetRetcodeShareUrlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 01FF8DD9-A09C-47A1-895A-B6E321******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The share URL of the application monitored by Browser Monitoring.
	//
	// example:
	//
	// http://arms-daily.console.aliyun.com:8080/shareapi/retcode.json?login_arms_t3h_token=XXXxxx&action=RetcodeAction&eventSubmitDoGetData=1
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRetcodeShareUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRetcodeShareUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRetcodeShareUrlResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetRetcodeShareUrlResponseBody) SetRequestId(v string) *GetRetcodeShareUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRetcodeShareUrlResponseBody) SetUrl(v string) *GetRetcodeShareUrlResponseBody {
	s.Url = &v
	return s
}

func (s *GetRetcodeShareUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
