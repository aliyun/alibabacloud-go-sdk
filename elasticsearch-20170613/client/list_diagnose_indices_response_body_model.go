// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDiagnoseIndicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDiagnoseIndicesResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListDiagnoseIndicesResponseBody
	GetResult() []*string
}

type ListDiagnoseIndicesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F05ED12E-140A-4ACB-B059-3A508A69F2E1
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDiagnoseIndicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnoseIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnoseIndicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDiagnoseIndicesResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListDiagnoseIndicesResponseBody) SetRequestId(v string) *ListDiagnoseIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnoseIndicesResponseBody) SetResult(v []*string) *ListDiagnoseIndicesResponseBody {
	s.Result = v
	return s
}

func (s *ListDiagnoseIndicesResponseBody) Validate() error {
	return dara.Validate(s)
}
