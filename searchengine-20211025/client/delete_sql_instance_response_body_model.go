// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSqlInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSqlInstanceResponseBody
	GetRequestId() *string
	SetResult(v *DeleteSqlInstanceResponseBodyResult) *DeleteSqlInstanceResponseBody
	GetResult() *DeleteSqlInstanceResponseBodyResult
}

type DeleteSqlInstanceResponseBody struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Response<Map<String, String>>
	Result *DeleteSqlInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteSqlInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSqlInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSqlInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSqlInstanceResponseBody) GetResult() *DeleteSqlInstanceResponseBodyResult {
	return s.Result
}

func (s *DeleteSqlInstanceResponseBody) SetRequestId(v string) *DeleteSqlInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSqlInstanceResponseBody) SetResult(v *DeleteSqlInstanceResponseBodyResult) *DeleteSqlInstanceResponseBody {
	s.Result = v
	return s
}

func (s *DeleteSqlInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteSqlInstanceResponseBodyResult struct {
	// id of request
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result map[string]*string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteSqlInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteSqlInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteSqlInstanceResponseBodyResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSqlInstanceResponseBodyResult) GetResult() map[string]*string {
	return s.Result
}

func (s *DeleteSqlInstanceResponseBodyResult) SetRequestId(v string) *DeleteSqlInstanceResponseBodyResult {
	s.RequestId = &v
	return s
}

func (s *DeleteSqlInstanceResponseBodyResult) SetResult(v map[string]*string) *DeleteSqlInstanceResponseBodyResult {
	s.Result = v
	return s
}

func (s *DeleteSqlInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
