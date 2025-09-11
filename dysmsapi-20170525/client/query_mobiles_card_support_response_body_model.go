// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMobilesCardSupportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMobilesCardSupportResponseBody
	GetCode() *string
	SetData(v *QueryMobilesCardSupportResponseBodyData) *QueryMobilesCardSupportResponseBody
	GetData() *QueryMobilesCardSupportResponseBodyData
	SetRequestId(v string) *QueryMobilesCardSupportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMobilesCardSupportResponseBody
	GetSuccess() *bool
}

type QueryMobilesCardSupportResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryMobilesCardSupportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 08C17DFE-2E10-54F4-BAFB-7180039CC217
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMobilesCardSupportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMobilesCardSupportResponseBody) GetData() *QueryMobilesCardSupportResponseBodyData {
	return s.Data
}

func (s *QueryMobilesCardSupportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMobilesCardSupportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMobilesCardSupportResponseBody) SetCode(v string) *QueryMobilesCardSupportResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetData(v *QueryMobilesCardSupportResponseBodyData) *QueryMobilesCardSupportResponseBody {
	s.Data = v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetRequestId(v string) *QueryMobilesCardSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) SetSuccess(v bool) *QueryMobilesCardSupportResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMobilesCardSupportResponseBodyData struct {
	// The list of returned results.
	QueryResult []*QueryMobilesCardSupportResponseBodyDataQueryResult `json:"QueryResult,omitempty" xml:"QueryResult,omitempty" type:"Repeated"`
}

func (s QueryMobilesCardSupportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBodyData) GetQueryResult() []*QueryMobilesCardSupportResponseBodyDataQueryResult {
	return s.QueryResult
}

func (s *QueryMobilesCardSupportResponseBodyData) SetQueryResult(v []*QueryMobilesCardSupportResponseBodyDataQueryResult) *QueryMobilesCardSupportResponseBodyData {
	s.QueryResult = v
	return s
}

func (s *QueryMobilesCardSupportResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryMobilesCardSupportResponseBodyDataQueryResult struct {
	// The mobile phone number.
	//
	// example:
	//
	// 1380000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Indicates whether the mobile phone number supports card messages. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
}

func (s QueryMobilesCardSupportResponseBodyDataQueryResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportResponseBodyDataQueryResult) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) GetMobile() *string {
	return s.Mobile
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) GetSupport() *bool {
	return s.Support
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) SetMobile(v string) *QueryMobilesCardSupportResponseBodyDataQueryResult {
	s.Mobile = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) SetSupport(v bool) *QueryMobilesCardSupportResponseBodyDataQueryResult {
	s.Support = &v
	return s
}

func (s *QueryMobilesCardSupportResponseBodyDataQueryResult) Validate() error {
	return dara.Validate(s)
}
