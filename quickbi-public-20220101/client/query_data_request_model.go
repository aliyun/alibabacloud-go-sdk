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
	// The API ID in [DataService Studio](https://help.aliyun.com/document_detail/144980.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// f4cc43bc3***
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// Filter conditions as a JSON map string. Each key is a request parameter name, and each value is the parameter value.
	//
	// **Note:**
	//
	// - If the operator of a request parameter is set to **Enumeration Filter**, the value can contain multiple values. In this case, the value must be in the format of a JSON list. For example: `area=["East China","North China","South China"]`
	//
	// - For dates, use the following formats based on the date type:
	//
	//   - Year: 2019
	//
	//   - Quarter: 2019Q1
	//
	//   - Month: 201901 (with a leading zero)
	//
	//   - Week: 2019-52
	//
	//   - Day: 20190101
	//
	//   - Hour: 14:00:00 (minutes and seconds are 00)
	//
	//   - Minute: 14:12:00 (seconds are 00)
	//
	//   - Second: 14:34:34
	//
	// example:
	//
	// { "area": ["test", "test"],  "shopping_date": "2019Q1",  }
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// A JSON array of field names to return.
	//
	// example:
	//
	// ["area", "city", "price", "date"]
	ReturnFields *string `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty"`
	// The Quick BI user ID. Obtain this value from [QueryUserInfoByAccount](https://next.api.aliyun.com/document/quickbi-public/2022-01-01/QueryUserInfoByAccount).
	//
	// > Specifies the user identity for DataService Studio, used with row-level and column-level permission configurations.
	//
	// 	Notice:
	//
	// If omitted, empty, or null, defaults to the Quick BI organization owner\\"s user ID.
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
