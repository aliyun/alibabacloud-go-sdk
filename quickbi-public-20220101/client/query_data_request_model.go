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
	// The ID of the API in Data Service. For more information, see [Data Service](https://help.aliyun.com/document_detail/144980.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// f4cc43bc3***
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The query conditions for Data Service. Specify the conditions as key-value pairs in a map-type string. The key is the request parameter name, and the value is the request parameter value. Keys and values must appear in pairs.
	//
	// **Note:**
	//
	// - When the operator of the request parameter is set to **enumeration filter**, the value can contain multiple values. In this case, the value must be a JSON-formatted list. Example: `area=["East China","North China","South China"]`
	//
	// - For dates, different input formats are provided based on the type:
	//
	//     - Year: 2019
	//
	//     - Quarter: 2019Q1
	//
	//     - Month: 201901 (with a leading zero)
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
	// { "area": ["test", "test"],  "shopping_date": "2019Q1",  }
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The list of names of the return parameters, which is a list-type string.
	//
	// example:
	//
	// ["area", "city", "price", "date"]
	ReturnFields *string `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty"`
	// The user ID in Quick BI. For information about how to obtain the user ID, see [Query user information by account](https://next.api.aliyun.com/document/quickbi-public/2022-01-01/QueryUserInfoByAccount).
	//
	// > This parameter specifies the identity of the user who uses Data Service. It can be used in combination with the row-level and column-level permission configurations of the dataset.
	//
	//
	//
	// 	Notice: If this parameter is not specified, is left empty, or is set to null, the user ID of the owner of the current Quick BI organization is used by default.</notice>
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
