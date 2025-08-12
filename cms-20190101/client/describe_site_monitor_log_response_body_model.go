// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSiteMonitorLogResponseBody
	GetCode() *string
	SetData(v string) *DescribeSiteMonitorLogResponseBody
	GetData() *string
	SetMessage(v string) *DescribeSiteMonitorLogResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeSiteMonitorLogResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSiteMonitorLogResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSiteMonitorLogResponseBody
	GetSuccess() *string
}

type DescribeSiteMonitorLogResponseBody struct {
	// The HTTP status code.
	//
	// **
	//
	// **Description*	- The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logs of the instant test tasks.
	//
	// example:
	//
	// [{\\\\"redirectCount\\\\":0.0,\\\\"SSLConnectTime\\\\":0.0,\\\\"pingDetail\\\\":\\\\"\\\\",\\\\"HTTPConnectTime\\\\":0.0,\\\\"isp\\\\":\\\\"465\\\\",\\\\"errorCode\\\\":611,\\\\"ispCN\\\\":\\\\"Alibaba\\\\",\\\\"resolution\\\\":\\\\"\\\\",\\\\"areaEN\\\\":\\\\"HuaBei\\\\",\\\\"taskEndTimestamp\\\\":1638422475687,\\\\"targetIspEN\\\\":\\\\"\\\\",\\\\"TotalTime\\\\":1.0,\\\\"taskStartTimestamp\\\\":1638422474389,\\\\"countryCN\\\\":\\\\"China\\\\",\\\\"provinceEN\\\\":\\\\"Beijing\\\\",\\\\"countryEN\\\\":\\\\"China\\\\",\\\\"targetCityEN\\\\":\\\\"\\\\",\\\\"curlConnectTime\\\\":0.0,\\\\"ips\\\\":\\\\"\\\\",\\\\"route\\\\":\\\\"\\\\",\\\\"tcpConnectTime\\\\":0.0,\\\\"cityEN\\\\":\\\\"Beijing\\\\",\\\\"HTTPDownloadSpeed\\\\":0.0,\\\\"HTTPDownloadTime\\\\":0.0,\\\\"HTTPResponseCode\\\\":0.0,\\\\"areaCN\\\\":\\\\"North China\\\\",\\\\"city\\\\":\\\\"546\\\\",\\\\"expection\\\\":\\\\"\\\\",\\\\"suorceIp\\\\":\\\\"192.168.XX.XX \\\\",\\\\"ispEN\\\\":\\\\"Alibaba\\\\",\\\\"HTTPDNSTime\\\\":1.0,\\\\"targetIsp\\\\":\\\\"\\\\",\\\\"curlStarttransferTime\\\\":0.0,\\\\"provinceCN\\\\":\\\\"Beijing\\\\",\\\\"timestamp\\\\":1638422474000,\\\\"redirectTime\\\\":0.0,\\\\"targetCity\\\\":\\\\"\\\\", \\\\"expect\\\\":\\\\"\\\\",\\\\"HTTPDownloadSize\\\\":0.0,\\\\"localDns\\\\":\\\\"192.168.XX.XX\\\\",\\\\"cityCN\\\\":\\\\"Beijing\\\\",\\\\"taskId\\\\":\\\\"afa5c3ce-f944-4363-9edb-ce919a29\\*\\*\\*\\*\\\\"}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67F646FA-ED8A-58C2-B461-451DB52C8B14
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSiteMonitorLogResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeSiteMonitorLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSiteMonitorLogResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSiteMonitorLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteMonitorLogResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSiteMonitorLogResponseBody) SetCode(v string) *DescribeSiteMonitorLogResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorLogResponseBody) SetData(v string) *DescribeSiteMonitorLogResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSiteMonitorLogResponseBody) SetMessage(v string) *DescribeSiteMonitorLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorLogResponseBody) SetNextToken(v string) *DescribeSiteMonitorLogResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSiteMonitorLogResponseBody) SetRequestId(v string) *DescribeSiteMonitorLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorLogResponseBody) SetSuccess(v string) *DescribeSiteMonitorLogResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorLogResponseBody) Validate() error {
	return dara.Validate(s)
}
