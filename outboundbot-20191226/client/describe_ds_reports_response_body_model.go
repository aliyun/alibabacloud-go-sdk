// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDsReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDsReportsResponseBody
	GetCode() *string
	SetDsReports(v string) *DescribeDsReportsResponseBody
	GetDsReports() *string
	SetHttpStatusCode(v int32) *DescribeDsReportsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeDsReportsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDsReportsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDsReportsResponseBody
	GetSuccess() *bool
}

type DescribeDsReportsResponseBody struct {
	// The API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dialog flow information in the format: key:{value:aaa}}, where aaa indicates the number of occurrences.
	//
	// example:
	//
	// "{\\"Label1\\":{\\"LabelValue1\\":2},\\"Label2\\":{\\"LabelValue1\\":2},\\"Label3\\":{\\"LabelValue1\\":1,\\"LabelValue2\\":1}}"
	DsReports *string `json:"DsReports,omitempty" xml:"DsReports,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The API response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1364f208-982d-4d0c-89aa-d56e22b47589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDsReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDsReportsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDsReportsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDsReportsResponseBody) GetDsReports() *string {
	return s.DsReports
}

func (s *DescribeDsReportsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDsReportsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDsReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDsReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDsReportsResponseBody) SetCode(v string) *DescribeDsReportsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDsReportsResponseBody) SetDsReports(v string) *DescribeDsReportsResponseBody {
	s.DsReports = &v
	return s
}

func (s *DescribeDsReportsResponseBody) SetHttpStatusCode(v int32) *DescribeDsReportsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDsReportsResponseBody) SetMessage(v string) *DescribeDsReportsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDsReportsResponseBody) SetRequestId(v string) *DescribeDsReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDsReportsResponseBody) SetSuccess(v bool) *DescribeDsReportsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDsReportsResponseBody) Validate() error {
	return dara.Validate(s)
}
