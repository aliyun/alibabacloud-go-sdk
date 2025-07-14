// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *QueryDataRequest
	GetApiId() *string
	SetConditions(v string) *QueryDataRequest
	GetConditions() *string
	SetReturnFields(v string) *QueryDataRequest
	GetReturnFields() *string
	SetUserId(v string) *QueryDataRequest
	GetUserId() *string
}

type QueryDataRequest struct {
	// The API ID in the data service. For more information, see: [Data Service](https://help.aliyun.com/document_detail/144980.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// f4cc43bc3***
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The query conditions for the data service, passed in as Key and Value pairs. A map-type string. Here, Key is the name of the request parameter, and Value is the value of the request parameter. Key and Value must appear in pairs.
	//
	// **Note:**
	//
	// - When the operator of the request parameter is set to **enumeration filtering**, the value can contain multiple values, and the format of the value should be a JSON-formatted List. For example: `area=["East China","North China","South China"]`
	//
	// - For dates, different formats are provided based on the type:
	//
	//     - Year: 2019
	//
	//     - Quarter: 2019Q1
	//
	//     - Month: 201901 (with leading zero)
	//
	//
	//
	//     - Week: 2019-52
	//
	//     - Day: 20190101
	//
	//     - Hour: 14:00:00 (minutes and seconds are 00)
	//
	//
	//
	//     - Minute: 14:12:00 (seconds are 00)
	//
	//     - Second: 14:34:34
	//
	// example:
	//
	// test
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// A list of return parameter names, in a List-type string.
	//
	// example:
	//
	// ["area", "city", "price", "date"]
	ReturnFields *string `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty"`
	// The userId in Quick BI. For how to obtain the userId, see: [Query User Information by Account Interface](https://next.api.aliyun.com/document/quickbi-public/2022-01-01/QueryUserInfoByAccount)
	//
	// > This parameter is used to specify the identity of the person using the data service, which can be used in conjunction with the row and column permission configurations of the dataset.
	//
	//
	//
	// 	Notice: If the parameter is not passed, an empty string is passed, or null is passed, the default userId will be the owner of the current Quick BI organization.</notice>
	//
	// example:
	//
	// b5d8fd9348cc4327****afb604
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDataRequest) GetApiId() *string {
	return s.ApiId
}

func (s *QueryDataRequest) GetConditions() *string {
	return s.Conditions
}

func (s *QueryDataRequest) GetReturnFields() *string {
	return s.ReturnFields
}

func (s *QueryDataRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryDataRequest) SetApiId(v string) *QueryDataRequest {
	s.ApiId = &v
	return s
}

func (s *QueryDataRequest) SetConditions(v string) *QueryDataRequest {
	s.Conditions = &v
	return s
}

func (s *QueryDataRequest) SetReturnFields(v string) *QueryDataRequest {
	s.ReturnFields = &v
	return s
}

func (s *QueryDataRequest) SetUserId(v string) *QueryDataRequest {
	s.UserId = &v
	return s
}

func (s *QueryDataRequest) Validate() error {
	return dara.Validate(s)
}
