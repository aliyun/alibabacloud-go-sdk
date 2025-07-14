// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSourceBloodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DataSourceBloodResponseBody
	GetRequestId() *string
	SetResult(v []*string) *DataSourceBloodResponseBody
	GetResult() []*string
	SetSuccess(v bool) *DataSourceBloodResponseBody
	GetSuccess() *bool
}

type DataSourceBloodResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 46e537a5****,3dadsu****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Array of dataset IDs.
	Result []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DataSourceBloodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DataSourceBloodResponseBody) GoString() string {
	return s.String()
}

func (s *DataSourceBloodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DataSourceBloodResponseBody) GetResult() []*string {
	return s.Result
}

func (s *DataSourceBloodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DataSourceBloodResponseBody) SetRequestId(v string) *DataSourceBloodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DataSourceBloodResponseBody) SetResult(v []*string) *DataSourceBloodResponseBody {
	s.Result = v
	return s
}

func (s *DataSourceBloodResponseBody) SetSuccess(v bool) *DataSourceBloodResponseBody {
	s.Success = &v
	return s
}

func (s *DataSourceBloodResponseBody) Validate() error {
	return dara.Validate(s)
}
