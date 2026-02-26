// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDisburseResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfirmDisburseResult
	GetRequestId() *string
	SetResult(v string) *ConfirmDisburseResult
	GetResult() *string
}

type ConfirmDisburseResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ConfirmDisburseResult) String() string {
	return dara.Prettify(s)
}

func (s ConfirmDisburseResult) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmDisburseResult) GetResult() *string {
	return s.Result
}

func (s *ConfirmDisburseResult) SetRequestId(v string) *ConfirmDisburseResult {
	s.RequestId = &v
	return s
}

func (s *ConfirmDisburseResult) SetResult(v string) *ConfirmDisburseResult {
	s.Result = &v
	return s
}

func (s *ConfirmDisburseResult) Validate() error {
	return dara.Validate(s)
}
