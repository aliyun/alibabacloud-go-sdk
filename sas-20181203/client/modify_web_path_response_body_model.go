// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebPathResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyWebPathResponseBody
	GetSuccess() *bool
}

type ModifyWebPathResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 20623245-7E5E-52CA-9640-7502F119****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyWebPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPathResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebPathResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyWebPathResponseBody) SetRequestId(v string) *ModifyWebPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebPathResponseBody) SetSuccess(v bool) *ModifyWebPathResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyWebPathResponseBody) Validate() error {
	return dara.Validate(s)
}
