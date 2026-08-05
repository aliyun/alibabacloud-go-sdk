// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTablesResponseBody
	GetRequestId() *string
	SetResult(v []*string) *GetTablesResponseBody
	GetResult() []*string
}

type GetTablesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 33E4F0CA-F766-5803-B11C-70DC57A5A6E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTablesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTablesResponseBody) GetResult() []*string {
	return s.Result
}

func (s *GetTablesResponseBody) SetRequestId(v string) *GetTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTablesResponseBody) SetResult(v []*string) *GetTablesResponseBody {
	s.Result = v
	return s
}

func (s *GetTablesResponseBody) Validate() error {
	return dara.Validate(s)
}
