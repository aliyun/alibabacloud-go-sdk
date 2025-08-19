// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppResponseBody
	GetRequestId() *string
	SetResult(v *DeleteAppResponseBodyResult) *DeleteAppResponseBody
	GetResult() *DeleteAppResponseBodyResult
}

type DeleteAppResponseBody struct {
	// example:
	//
	// 2C5DAA30-****-5181-9B87-9D6181016197
	RequestId *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *DeleteAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DeleteAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppResponseBody) GetResult() *DeleteAppResponseBodyResult {
	return s.Result
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetResult(v *DeleteAppResponseBodyResult) *DeleteAppResponseBody {
	s.Result = v
	return s
}

func (s *DeleteAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteAppResponseBodyResult struct {
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s DeleteAppResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAppResponseBodyResult) SetInstanceId(v string) *DeleteAppResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DeleteAppResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
