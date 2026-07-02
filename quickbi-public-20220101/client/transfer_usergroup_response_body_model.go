// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferUsergroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferUsergroupResponseBody
	GetRequestId() *string
	SetResult(v bool) *TransferUsergroupResponseBody
	GetResult() *bool
	SetSuccess(v bool) *TransferUsergroupResponseBody
	GetSuccess() *bool
}

type TransferUsergroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC4E1E63-B337-44F8-8C22-6F00DF67E2C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the migration was successful.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferUsergroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferUsergroupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferUsergroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferUsergroupResponseBody) GetResult() *bool {
	return s.Result
}

func (s *TransferUsergroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferUsergroupResponseBody) SetRequestId(v string) *TransferUsergroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferUsergroupResponseBody) SetResult(v bool) *TransferUsergroupResponseBody {
	s.Result = &v
	return s
}

func (s *TransferUsergroupResponseBody) SetSuccess(v bool) *TransferUsergroupResponseBody {
	s.Success = &v
	return s
}

func (s *TransferUsergroupResponseBody) Validate() error {
	return dara.Validate(s)
}
