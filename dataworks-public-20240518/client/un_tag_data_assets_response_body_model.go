// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagDataAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnTagDataAssetsResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UnTagDataAssetsResponseBody
	GetSuccess() *string
}

type UnTagDataAssetsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8754EE08-4AA2-5F77-ADD7-754DBBDA9F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnTagDataAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnTagDataAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagDataAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnTagDataAssetsResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UnTagDataAssetsResponseBody) SetRequestId(v string) *UnTagDataAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnTagDataAssetsResponseBody) SetSuccess(v string) *UnTagDataAssetsResponseBody {
	s.Success = &v
	return s
}

func (s *UnTagDataAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}
