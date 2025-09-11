// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDataSourceResponseBody
	GetRequestId() *string
	SetResult(v string) *AddDataSourceResponseBody
	GetResult() *string
	SetSuccess(v bool) *AddDataSourceResponseBody
	GetSuccess() *bool
}

type AddDataSourceResponseBody struct {
	// example:
	//
	// D787E1********DF8D885
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 472b241b-c651-****-****-ee719d6faf45
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataSourceResponseBody) GetResult() *string {
	return s.Result
}

func (s *AddDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDataSourceResponseBody) SetRequestId(v string) *AddDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataSourceResponseBody) SetResult(v string) *AddDataSourceResponseBody {
	s.Result = &v
	return s
}

func (s *AddDataSourceResponseBody) SetSuccess(v bool) *AddDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *AddDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
