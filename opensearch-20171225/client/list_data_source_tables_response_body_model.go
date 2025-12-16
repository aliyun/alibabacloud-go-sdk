// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataSourceTablesResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListDataSourceTablesResponseBody
	GetResult() []*string
}

type ListDataSourceTablesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceTablesResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListDataSourceTablesResponseBody) SetRequestId(v string) *ListDataSourceTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTablesResponseBody) SetResult(v []*string) *ListDataSourceTablesResponseBody {
	s.Result = v
	return s
}

func (s *ListDataSourceTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
