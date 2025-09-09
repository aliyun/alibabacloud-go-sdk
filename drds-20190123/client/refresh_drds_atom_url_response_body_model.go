// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDrdsAtomUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshDrdsAtomUrlResponseBody
	GetRequestId() *string
	SetResult(v bool) *RefreshDrdsAtomUrlResponseBody
	GetResult() *bool
	SetSuccess(v bool) *RefreshDrdsAtomUrlResponseBody
	GetSuccess() *bool
}

type RefreshDrdsAtomUrlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B12FC174-D5CE-4A6E-83C1-0F8F86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the connection after refresh was successful.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshDrdsAtomUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshDrdsAtomUrlResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDrdsAtomUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshDrdsAtomUrlResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RefreshDrdsAtomUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefreshDrdsAtomUrlResponseBody) SetRequestId(v string) *RefreshDrdsAtomUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponseBody) SetResult(v bool) *RefreshDrdsAtomUrlResponseBody {
	s.Result = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponseBody) SetSuccess(v bool) *RefreshDrdsAtomUrlResponseBody {
	s.Success = &v
	return s
}

func (s *RefreshDrdsAtomUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
