// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSiteMonitorDataResponseBody
	GetCode() *string
	SetData(v string) *DescribeSiteMonitorDataResponseBody
	GetData() *string
	SetMessage(v string) *DescribeSiteMonitorDataResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeSiteMonitorDataResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSiteMonitorDataResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSiteMonitorDataResponseBody
	GetSuccess() *string
}

type DescribeSiteMonitorDataResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring data.
	//
	// example:
	//
	// [{"Maximum":247,"Mimimum":61,"Average":154,"userId":"127067667954****","taskId":"49f7b317-7645-4cc9-94fd-ea42e522****","timestamp":1551581760000}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// ea42e5220930ea42e522****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3febb181-0d98-4af9-8b04-7faf36b048b9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true: The request was successful. false: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSiteMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSiteMonitorDataResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeSiteMonitorDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSiteMonitorDataResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSiteMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSiteMonitorDataResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSiteMonitorDataResponseBody) SetCode(v string) *DescribeSiteMonitorDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetData(v string) *DescribeSiteMonitorDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetMessage(v string) *DescribeSiteMonitorDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetNextToken(v string) *DescribeSiteMonitorDataResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetRequestId(v string) *DescribeSiteMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) SetSuccess(v string) *DescribeSiteMonitorDataResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSiteMonitorDataResponseBody) Validate() error {
	return dara.Validate(s)
}
