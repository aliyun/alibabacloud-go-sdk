// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *QueryDataServiceRequest
	GetApiId() *string
	SetConditions(v string) *QueryDataServiceRequest
	GetConditions() *string
	SetReturnFields(v string) *QueryDataServiceRequest
	GetReturnFields() *string
}

type QueryDataServiceRequest struct {
	// The API ID in the data service. For more information, see [Data Service](https://help.aliyun.com/document_detail/144980.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// f4cc43bc3***
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The query conditions for the data service, passed in as Key-Value pairs. This is a map-type string. Here, Key is the name of the request parameter, and Value is the value of the request parameter. Keys and Values must appear in pairs.
	//
	// **Note:**
	//
	// - When the operator of the request parameter is set to **enumeration filter**, the value can contain multiple values. In this case, the format of the value is a JSON list. For example: `area=["East China","North China","South China"]`
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
	// { "area": ["华东", "华北"],  "shopping_date": "2019Q1",  }
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// A list of parameter names to be returned, as a List-type string.
	//
	// example:
	//
	// ["area", "city", "price", "date"]
	ReturnFields *string `json:"ReturnFields,omitempty" xml:"ReturnFields,omitempty"`
}

func (s QueryDataServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceRequest) GoString() string {
	return s.String()
}

func (s *QueryDataServiceRequest) GetApiId() *string {
	return s.ApiId
}

func (s *QueryDataServiceRequest) GetConditions() *string {
	return s.Conditions
}

func (s *QueryDataServiceRequest) GetReturnFields() *string {
	return s.ReturnFields
}

func (s *QueryDataServiceRequest) SetApiId(v string) *QueryDataServiceRequest {
	s.ApiId = &v
	return s
}

func (s *QueryDataServiceRequest) SetConditions(v string) *QueryDataServiceRequest {
	s.Conditions = &v
	return s
}

func (s *QueryDataServiceRequest) SetReturnFields(v string) *QueryDataServiceRequest {
	s.ReturnFields = &v
	return s
}

func (s *QueryDataServiceRequest) Validate() error {
	return dara.Validate(s)
}
